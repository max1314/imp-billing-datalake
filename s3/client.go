// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package s3

import (
	"bytes"
	"context"
	"encoding/hex"
	"io"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
	"time"

	blob "github.com/max1314/imp-billing-datalake"
	"github.com/max1314/imp-billing-datalake/internal/buffer"
	"github.com/max1314/imp-billing-datalake/internal/errors"
	"github.com/max1314/imp-billing-datalake/internal/extclients/aws/awserr"
	"github.com/max1314/imp-billing-datalake/internal/reporting"
	"github.com/max1314/imp-billing-datalake/internal/sharedflags"
	// Empty depedencies so go-dep doesn't complain.
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	fSignedUrlExpirationTime = sharedflags.Set.Duration(
		"s3_signed_url_expiration_time",
		20*time.Minute,
		"The expiration time for signed urls for the storage API in minutes")
	fCredsKey = sharedflags.Set.String(
		"s3_credentials_key_id", "",
		"S3 credentials key id.")
	fCredsKeyFile = sharedflags.Set.String(
		"s3_credentials_key_id_file", "",
		"Path to S3 credentials file containing the AWS_KEY_ID")
	fCredsSecretFile = sharedflags.Set.String(
		"s3_credentials_key_secret_file", "",
		"Path to S3 credentials secret file containing the secret key.")
	fCredsSecret = sharedflags.Set.String(
		"s3_credentials_secret", "",
		"S3 sercret key (by value) to be supplied from Vault")
	fEndpoint = sharedflags.Set.String(
		"s3_endpoint", "s3.eu-west-2.amazonaws.com",
		"S3 endpoint")
	fZone = sharedflags.Set.String(
		"s3_region", "eu-west-2",
		"S3 region")
	fS3ForcePathStyle = sharedflags.Set.Bool(
		"s3_force_path_style", false,
		"Have s3 use path style names (s3.com/mybucket/file) instead of virtual host (mybucket.s3.com/file).")
)

var (
	schemeRegex = regexp.MustCompile(`^https?://`)
)

// Client abstraction over s3 client to provide upload and download signed URLs for limited time.
type dataLake struct {
	ctx        context.Context
	svc        *s3.S3
	bucketName string
}

// You probably want to use imp-billing-datalake/util.NewBucketFromFlag instead for easier support for multiple clouds.
func NewDataLake(ctx context.Context, bucketName string) (blob.DataLake, error) {
	config := &aws.Config{}

	if *fCredsKey != "" || *fCredsKeyFile != "" {
		keyId := *fCredsKey
		if *fCredsKeyFile != "" {
			b, err := ioutil.ReadFile(*fCredsKeyFile)
			if err != nil {
				return nil, errors.Wrapf(err, "could not read credentials key file from path %v", *fCredsKeyFile)
			}
			keyId = strings.TrimSpace(string(b))
		}

		var secret string
		if *fCredsSecret != "" {
			secret = *fCredsSecret
		} else {
			if *fCredsSecretFile == "" {
				return nil, errors.New(nil, errors.InvalidArgument, "you must set both the key id and the file with the secret key")
			}
			secretData, err := ioutil.ReadFile(*fCredsSecretFile)
			if err != nil {
				return nil, errors.Wrap(err, "couldn't read file with secret key")
			}
			secret = strings.TrimSpace(string(secretData))
		}

		config.Credentials = credentials.NewStaticCredentials(keyId, secret, "")
	}

	endpoint := *fEndpoint
	if !schemeRegex.MatchString(endpoint) {
		endpoint = "https://" + endpoint
	}
	config.Endpoint = aws.String(endpoint)
	config.S3ForcePathStyle = aws.Bool(*fS3ForcePathStyle)

	if *fZone != "" {
		config.Region = aws.String(*fZone)
	}

	ses, err := session.NewSession(config)
	if err != nil {
		return nil, errors.Wrap(awserr.RemapError(err), "couldn't access aws services")
	}

	return &dataLake{
		ctx: ctx,
		svc: s3.New(
			ses,
			config,
		),
		bucketName: bucketName,
	}, nil
}

func (b *dataLake) Exists(path string) (exists bool, err error) {
	defer reporting.Track(reporting.Exists, blob.S3, &err)()

	// AWS S3 behaves weirdly when calling `HeadObject` and the object doesn't exist. Instead of
	// returning a `NoSuchKey` error, it returns a HTTP 404, which is annoying to deal with. Using
	// `GetObject` instead returns `NoSuchKey` properly on error, but gets the entire data for the
	// file on success. As such, the best option here is to see if an object name is present on the
	// list. This way we also avoid losing s3 read-after-write strong consistency for the first write
	// when checking for an object that doesn't exist (for more info,
	// https://docs.aws.amazon.com/AmazonS3/latest/dev/Introduction.html#API)
	r, err := b.svc.ListObjectsV2WithContext(b.ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(b.bucketName),
		Prefix: aws.String(path),
	})

	if err != nil {
		return false, errors.Wrapf(awserr.RemapError(err), "failed to check if object %s exists", path)
	}

	for _, content := range r.Contents {
		if *content.Key == path {
			return true, nil
		}
	}

	return false, nil
}

func (b *dataLake) NewReader(path string) (rc io.ReadCloser, err error) {
	defer reporting.Track(reporting.NewReader, blob.S3, &err)()

	r, err := b.svc.GetObjectWithContext(b.ctx, &s3.GetObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, errors.Wrapf(awserr.RemapError(err), "error getting the object (%s) from the S3 compatible bucket: %s", path, b.bucketName)
	}

	return r.Body, nil
}

func (b *dataLake) NewWriter(path string) io.WriteCloser {
	attrs := &blob.ObjectAttributes{}
	ctx := context.Background()
	return buffer.NewClosingBuffer(func(data []byte) (err error) {
		defer reporting.Track(reporting.NewWriter, blob.S3, &err)()

		put := &s3.PutObjectInput{
			Body:                 bytes.NewReader(data),
			Bucket:               aws.String(b.bucketName),
			Key:                  aws.String(path),
			ServerSideEncryption: aws.String(s3.ServerSideEncryptionAes256),
		}

		if attrs != nil {
			if attrs.CacheControl != "" {
				put.SetCacheControl(attrs.CacheControl)
			}
			if attrs.ContentEncoding != "" {
				put.SetContentEncoding(attrs.ContentEncoding)
			}
			if attrs.ContentDisposition != "" {
				put.SetContentEncoding(attrs.ContentDisposition)
			}
			if attrs.ContentLanguage != "" {
				put.SetContentLanguage(attrs.ContentLanguage)
			}
			if attrs.ContentType != "" {
				put.SetContentType(attrs.ContentType)
			}

			if attrs.Metadata != nil {
				metadata := make(map[string]*string)
				for k, v := range attrs.Metadata {
					// Turns out that when you range through a map, `k, v` are reused variables that constantly
					// get a copy of the current "iterated" values, so taking &v directly will always return the
					// same address, which eventually points to the last iterated value. To force remembering
					// addresses to all values, we have to copy them in a new variable.
					copy := v
					metadata[k] = &copy
				}

				put.SetMetadata(metadata)
			}
		}

		_, err = b.svc.PutObjectWithContext(ctx, put)

		if err != nil {
			return errors.Wrap(awserr.RemapError(err), "error putting the object on the S3 compatible server")
		}

		return nil
	})
}

func (b *dataLake) Delete(path string) (err error) {
	defer reporting.Track(reporting.Delete, blob.S3, &err)()
	_, err = b.svc.DeleteObjectWithContext(b.ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(path),
	})

	if err != nil {
		return errors.Wrap(awserr.RemapError(err), "failed to delete S3 object")
	}
	return nil
}

func (b *dataLake) Copy(srcPath string, destPath string) (err error) {
	defer reporting.Track(reporting.Copy, blob.S3, &err)()
	_, err = b.svc.CopyObjectWithContext(b.ctx, &s3.CopyObjectInput{
		Bucket:               aws.String(b.bucketName),
		CopySource:           aws.String(path.Join(b.bucketName, srcPath)),
		Key:                  aws.String(destPath),
		ServerSideEncryption: aws.String(s3.ServerSideEncryptionAes256),
	})

	if err != nil {
		return errors.Wrap(awserr.RemapError(err), "failed to copy S3 object")
	}

	return nil
}

func (b *dataLake) List(dirPath string) (dirs []string, objects []string, err error) {
	defer reporting.Track(reporting.GetDirContent, blob.S3, &err)()
	r, err := b.svc.ListObjectsV2WithContext(context.Background(), &s3.ListObjectsV2Input{
		Bucket:    aws.String(b.bucketName),
		Delimiter: aws.String("/"),
		Prefix:    aws.String(dirPath),
	})

	if err != nil {
		return nil, nil, errors.Wrap(awserr.RemapError(err), "couldn't list objects at s3 compatible storage")
	}

	for _, pref := range r.CommonPrefixes {
		dirs = append(dirs, *pref.Prefix)
	}

	for _, obj := range r.Contents {
		objects = append(objects, *obj.Key)
	}

	return dirs, objects, nil
}

func toAttributes(s3Attr *s3.HeadObjectOutput) *blob.ObjectAttributes {
	var attr *blob.ObjectAttributes = &blob.ObjectAttributes{}
	if s3Attr.CacheControl != nil {
		attr.CacheControl = *s3Attr.CacheControl
	}
	if s3Attr.ContentEncoding != nil {
		attr.ContentEncoding = *s3Attr.ContentEncoding
	}
	if s3Attr.ContentDisposition != nil {
		attr.ContentDisposition = *s3Attr.ContentDisposition
	}
	if s3Attr.ContentLanguage != nil {
		attr.ContentLanguage = *s3Attr.ContentLanguage
	}
	if s3Attr.ContentType != nil {
		attr.ContentType = *s3Attr.ContentType
	}
	if s3Attr.ContentLength != nil {
		attr.Size = *s3Attr.ContentLength
	}
	if s3Attr.ETag != nil {
		//NOTE: We drop the error here because upstream
		//we assume (based on code inspection) that if the
		//md5 is malformed the code will error elsewhere with
		//appropriate information and context
		hexBytes, _ := hex.DecodeString(strings.Trim(*s3Attr.ETag, "\""))
		attr.MD5 = hexBytes
	}
	if s3Attr.LastModified != nil {
		attr.LastModified = *s3Attr.LastModified
	}
	if s3Attr.Metadata != nil {
		attr.Metadata = map[string]string{}
		for k, v := range s3Attr.Metadata {
			attr.Metadata[k] = *v
		}
	}
	return attr
}


