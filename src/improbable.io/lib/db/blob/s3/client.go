// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package s3

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"improbable.io/lib/db/blob"
	"improbable.io/lib/db/blob/internal/buffer"
	"improbable.io/lib/db/blob/internal/reporting"
	"improbable.io/lib/errors"
	"improbable.io/lib/extclients/aws/awserr"
	"improbable.io/lib/sharedflags"
	"improbable.io/proto/improbable/platform"

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
type bucket struct {
	ctx        context.Context
	svc        *s3.S3
	bucketName string
}

// You probably want to use improbable.io/lib/db/blob/util.NewBucketFromFlag instead for easier support for multiple clouds.
func NewBucket(ctx context.Context, bucketName string) (blob.Bucket, error) {
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

	return &bucket{
		ctx: ctx,
		svc: s3.New(
			ses,
			config,
		),
		bucketName: bucketName,
	}, nil
}

func (b *bucket) Exists(path string) (exists bool, err error) {
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

func (b *bucket) GetSupportedFileStorageProvider() improbable_platform.StorageProvider {
	return improbable_platform.StorageProvider_S3
}

func (b *bucket) GetSignedUploadURL(path string, md5 string) (url string, err error) {
	defer reporting.Track(reporting.GetSignedUploadURL, blob.S3, &err)()

	if md5 == "" {
		err = errors.New(nil, errors.InvalidArgument, "a base64 encoded MD5 of the file to be uploaded must be included")
		return "", err
	}

	r, _ := b.svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:               aws.String(b.bucketName),
		ContentMD5:           aws.String(md5),
		Key:                  aws.String(path),
		ServerSideEncryption: aws.String(s3.ServerSideEncryptionAes256),
	})

	url, err = r.Presign(*fSignedUrlExpirationTime)
	if err != nil {
		return "", errors.Wrapf(awserr.RemapError(err), "failed to pre-sign S3 PutObjectRequest")
	}

	return url, nil
}

func (b *bucket) GetSignedResumableUploadURL(path string, md5 string) (url string, err error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by this blob client type")
}

func (b *bucket) GetSignedCreateMultipartUploadURL(path string) (url string, err error) {
	defer reporting.Track(reporting.GetSignedCreateMultipartUploadURL, blob.S3, &err)()

	r, _ := b.svc.CreateMultipartUploadRequest(&s3.CreateMultipartUploadInput{
		Bucket:               aws.String(b.bucketName),
		Key:                  aws.String(path),
		ServerSideEncryption: aws.String(s3.ServerSideEncryptionAes256),
	})

	url, err = r.Presign(*fSignedUrlExpirationTime)
	if err != nil {
		return "", errors.Wrap(awserr.RemapError(err), "failed to pre-sign S3 CreateMultipartUploadRequest")
	}

	return url, nil
}

func (b *bucket) GetSignedUploadPartURL(path string, md5 string, uploadId string, partNumber int64) (url string, err error) {
	defer reporting.Track(reporting.GetSignedUploadpartURL, blob.S3, &err)()

	r, _ := b.svc.UploadPartRequest(&s3.UploadPartInput{
		Bucket:     aws.String(b.bucketName),
		ContentMD5: aws.String(md5),
		Key:        aws.String(path),
		PartNumber: aws.Int64(partNumber),
		UploadId:   aws.String(uploadId),
	})

	url, err = r.Presign(*fSignedUrlExpirationTime)
	if err != nil {
		return "", errors.Wrap(awserr.RemapError(err), "failed to pre-sign S3 UploadPartRequest")
	}

	return url, nil
}

func (b *bucket) GetSignedCompleteMultipartUploadURL(path string, uploadId string) (url string, err error) {
	defer reporting.Track(reporting.GetSignedCompleteMultipartUploadURL, blob.S3, &err)()

	r, _ := b.svc.CompleteMultipartUploadRequest(&s3.CompleteMultipartUploadInput{
		Bucket:          aws.String(b.bucketName),
		Key:             aws.String(path),
		MultipartUpload: nil, // Setting this to nil seems fine
		UploadId:        aws.String(uploadId),
	})

	url, err = r.Presign(*fSignedUrlExpirationTime)
	if err != nil {
		return "", errors.Wrap(awserr.RemapError(err), "failed to pre-sign S3 CompleteMultipartUploadRequest")
	}

	return url, nil
}

func (b *bucket) GetSignedDownloadURL(path string, attr *blob.GetObjectAttributes) (url string, err error) {
	defer reporting.Track(reporting.GetSignedDownloadURL, blob.S3, &err)()
	s3GetObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(path),
		// We need to override the content type to "application/octet-stream" to ensure that Runtime will be able to correctly download artifacts from S3
		ResponseContentType: aws.String("application/octet-stream"),
	}
	if attr != nil {
		s3GetObjectInput.ResponseContentDisposition = aws.String(attr.ResponseContentDisposition)
	}
	r, _ := b.svc.GetObjectRequest(s3GetObjectInput)

	url, err = r.Presign(*fSignedUrlExpirationTime)
	if err != nil {
		return "", errors.Wrap(awserr.RemapError(err), "failed to presign S3 GET request")
	}

	return url, nil
}

func (b *bucket) NewReader(path string, _ *blob.ReaderOptions) (rc io.ReadCloser, err error) {
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

func (b *bucket) NewWriter(ctx context.Context, path string, attrs *blob.ObjectAttributes) io.WriteCloser {
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

func (b *bucket) UpdateMeta(objPath string, attrUpdate blob.ObjectAttributesUpdate) (err error) {
	defer reporting.Track(reporting.UpdateMeta, blob.S3, &err)()

	cpinput := &s3.CopyObjectInput{
		Bucket: aws.String(b.bucketName),
		// From AWS doc:
		// CopySource is the name of the source bucket and key name of the source object, separated
		// by a slash (/). Must be URL-encoded.
		CopySource: aws.String(path.Join(b.bucketName, objPath)),
		Key:        aws.String(objPath),

		CacheControl:         attrUpdate.CacheControl,
		ContentEncoding:      attrUpdate.ContentEncoding,
		ContentDisposition:   attrUpdate.ContentDisposition,
		ContentLanguage:      attrUpdate.ContentLanguage,
		ContentType:          attrUpdate.ContentType,
		ServerSideEncryption: aws.String(s3.ServerSideEncryptionAes256),
	}

	if attrUpdate.Metadata != nil {
		meta := map[string]*string{}
		for k, v := range attrUpdate.Metadata {
			// Turns out that when you range through a map, `k, v` are reused variables that constantly
			// get a copy of the current "iterated" values, so taking &v directly will always return the
			// same address, which eventually points to the last iterated value. To force remembering
			// addresses to all values, we have to copy them in a new variable.
			copy := v
			meta[k] = &copy
		}

		cpinput.Metadata = meta
	}

	// In order to update meta, we must set MetadataDirective = REPLACE
	// See https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectCOPY.html for details.
	cpinput.MetadataDirective = aws.String("REPLACE")

	_, err = b.svc.CopyObjectWithContext(b.ctx, cpinput)

	if err != nil {
		return errors.Wrap(awserr.RemapError(err), "failed to copy S3 object")
	}
	return nil
}

func (b *bucket) UpdateACL(path string, acl blob.ACL) (err error) {
	defer reporting.Track(reporting.UpdateACL, blob.S3, &err)()
	_, err = b.svc.PutObjectAclWithContext(b.ctx, &s3.PutObjectAclInput{
		ACL:    aws.String(aclToCannedACL(acl)),
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(path),
	})

	if err != nil {
		return errors.Wrap(awserr.RemapError(err), "error updating the object's ACL")
	}
	return nil
}

func (b *bucket) Delete(path string) (err error) {
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

func (b *bucket) GetPublicURL(path string) string {
	s3url, err := url.Parse(*b.svc.Config.Endpoint)
	if err == nil {
		s3url.Host = fmt.Sprintf("%s.%s", b.bucketName, s3url.Host)
		s3url.Path = path
		return s3url.String()
	}
	return ""
}

func (b *bucket) Stat(path string) (attrs *blob.ObjectAttributes, err error) {
	defer reporting.Track(reporting.Stat, blob.S3, &err)()
	r, err := b.svc.HeadObjectWithContext(b.ctx, &s3.HeadObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(path),
	})

	if err != nil {
		return nil, errors.Wrap(awserr.RemapError(err), "object doesn't exist or isn't readable")
	}
	return toAttributes(r), nil
}

func (b *bucket) Copy(srcPath string, destPath string) (err error) {
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

func (b *bucket) Name() string {
	return b.bucketName
}

func (b *bucket) GetDirContent(ctx context.Context, dirPath string) (dirs []string, objects []string, err error) {
	defer reporting.Track(reporting.GetDirContent, blob.S3, &err)()
	r, err := b.svc.ListObjectsV2WithContext(ctx, &s3.ListObjectsV2Input{
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

func (b *bucket) GetContentWithCommonPrefix(ctx context.Context, prefix string, maxResults int) (objects []string, err error) {
	defer reporting.Track(reporting.GetContentWithCommonPrefix, blob.S3, &err)()
	args := &s3.ListObjectsV2Input{
		Bucket: aws.String(b.bucketName),
		Prefix: aws.String(prefix),
	}

	if maxResults != 0 {
		args.MaxKeys = aws.Int64(int64(maxResults))
	}

	r, err := b.svc.ListObjectsV2WithContext(ctx, args)

	if err != nil {
		return nil, errors.Wrap(awserr.RemapError(err), "couldn't list objects at s3 compatible storage")
	}

	for _, obj := range r.Contents {
		objects = append(objects, *obj.Key)
	}
	return objects, nil
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

func aclToCannedACL(acl blob.ACL) string {
	switch acl {
	case blob.PUBLIC_READ:
		return "public-read"
	case blob.AUTHENTICATED_READ:
		return "authenticated-read"
	case blob.PRIVATE:
		fallthrough
	default:
		return "private"
	}
}
