// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	blob "github.com/improbable/imp-billing-datalake"
	"github.com/improbable/imp-billing-datalake/errors"
	"github.com/improbable/imp-billing-datalake/internal/buffer"
)

var (
	schemeRegex = regexp.MustCompile(`^https?://`)
)

// `true` is using `s3.com/mybucket/file` style and `false` is using `mybucket.s3.com/file`
const S3ForcePathStyle = true

// Client abstraction over s3 client to provide upload and download signed URLs for limited time.
type dataLake struct {
	ctx        context.Context
	svc        *s3.S3
	bucketName string
}

// You probably want to use imp-billing-datalake/util.NewBucketFromFlag instead for easier support for multiple clouds.
func NewDataLake(ctx context.Context, bucketName, s3Endpoint, region string, creds *credentials.Credentials) (blob.DataLake, error) {
	config := &aws.Config{}
	config.Credentials = creds

	if !schemeRegex.MatchString(s3Endpoint) {
		s3Endpoint = "https://" + s3Endpoint
	}
	config.Endpoint = aws.String(s3Endpoint)
	config.S3ForcePathStyle = aws.Bool(S3ForcePathStyle)

	if region != "" {
		config.Region = aws.String(region)
	}

	ses, err := session.NewSession(config)
	if err != nil {
		return nil, errors.New(err, "couldn't access aws services ")
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
		return false, errors.New(err, fmt.Sprintf("failed to check if object %s exists ", path))
	}

	for _, content := range r.Contents {
		if *content.Key == path {
			return true, nil
		}
	}

	return false, nil
}

func (b *dataLake) NewReader(path string) (rc io.ReadCloser, err error) {

	r, err := b.svc.GetObjectWithContext(b.ctx, &s3.GetObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, errors.New(err, fmt.Sprintf("error getting the object (%s) from the S3 compatible bucket: %s ", path, b.bucketName))
	}

	return r.Body, nil
}

func (b *dataLake) NewWriter(path string) io.WriteCloser {
	ctx := context.Background()
	return buffer.NewClosingBuffer(func(data []byte) (err error) {

		put := &s3.PutObjectInput{
			Body:                 bytes.NewReader(data),
			Bucket:               aws.String(b.bucketName),
			Key:                  aws.String(path),
			ServerSideEncryption: aws.String(s3.ServerSideEncryptionAes256),
		}

		_, err = b.svc.PutObjectWithContext(ctx, put)

		if err != nil {
			panic(err)
			//return errors.New(err,"error putting the object on the S3 compatible server")
		}

		return nil
	})
}

func (b *dataLake) Delete(path string) (err error) {
	_, err = b.svc.DeleteObjectWithContext(b.ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(path),
	})

	if err != nil {
		return errors.New(err, "failed to delete S3 object ")
	}
	return nil
}

func (b *dataLake) Copy(srcPath string, destPath string) (err error) {
	_, err = b.svc.CopyObjectWithContext(b.ctx, &s3.CopyObjectInput{
		Bucket:               aws.String(b.bucketName),
		CopySource:           aws.String(path.Join(b.bucketName, srcPath)),
		Key:                  aws.String(destPath),
		ServerSideEncryption: aws.String(s3.ServerSideEncryptionAes256),
	})

	if err != nil {
		return errors.New(err, "failed to copy S3 object ")
	}

	return nil
}

func (b *dataLake) List(dirPath string) (dirs []string, objects []string, err error) {
	r, err := b.svc.ListObjectsV2WithContext(context.Background(), &s3.ListObjectsV2Input{
		Bucket:    aws.String(b.bucketName),
		Delimiter: aws.String("/"),
		Prefix:    aws.String(dirPath),
	})

	if err != nil {
		return nil, nil, errors.New(err, "couldn't list objects at s3 compatible storage ")
	}

	for _, pref := range r.CommonPrefixes {
		dirs = append(dirs, *pref.Prefix)
	}

	for _, obj := range r.Contents {
		objects = append(objects, *obj.Key)
	}

	return dirs, objects, nil
}
