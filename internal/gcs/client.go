// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Contains abstraction over the Google storage API.

package gcs

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	blob "github.com/improbable/imp-billing-datalake"
	"github.com/improbable/imp-billing-datalake/errors"
)

const GCSReadCompressed=true

// NewDataLake creates a client with explicit values of creds and project
func NewDataLake(ctx context.Context, bucketName string, opts ...option.ClientOption) (blob.DataLake, error) {
	var err error

	bucket := &dataLake{bucketName: bucketName, ctx: ctx}
	bucket.client, err = storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	bucket.bucket = bucket.client.Bucket(bucketName)
	return bucket, nil
}

// NewDataLakeWithTokenSource creates a new GCS bucket from OIDC token source.
func NewDataLakeWithTokenSource(ctx context.Context, bucketName string, source oauth2.TokenSource) (blob.DataLake, error) {
	client, err := storage.NewClient(ctx, option.WithTokenSource(source))
	if err != nil {
		return nil, err
	}

	return &dataLake{
		bucket:     client.Bucket(bucketName),
		ctx:        ctx,
		client:     client,
		bucketName: bucketName,
	}, nil
}

type dataLake struct {
	ctx               context.Context
	client            *storage.Client
	bucket            *storage.BucketHandle
	bucketName        string
	urlSigningAccount string
	urlSigningKey     []byte
}

func (b *dataLake) NewReader(path string) (r io.ReadCloser, err error) {

	object := b.bucket.Object(path)
	object = object.ReadCompressed(GCSReadCompressed)

	reader, err := object.NewReader(b.ctx)
	if err != nil {
		panic(err)
		return nil, errors.New(err, fmt.Sprintf("Could not create reader for %s ", path))
	}
	return reader, nil
}

func (b *dataLake) NewWriter(path string) io.WriteCloser {
	w := b.bucket.Object(path).NewWriter(b.ctx)
	return w
}


func (b *dataLake) Exists(path string) (exists bool, err error) {
	_, err = b.bucket.Object(path).Attrs(b.ctx)
	if err != nil {
		return false, errors.New(err, fmt.Sprintf("could not check if %s exists ", path))
	}
	return true, nil
}

func (b *dataLake) Delete(path string) (err error) {
	err = b.bucket.Object(path).Delete(b.ctx)
	if err != nil {
		return errors.New(err, fmt.Sprintf("could not delete %s", path))
	}
	return nil
}

// Copy calls google api to do a copy
func (b *dataLake) Copy(srcName string, destName string) (err error) {

	src := b.bucket.Object(srcName)
	dst := b.bucket.Object(destName)

	copier := dst.CopierFrom(src)

	_, err = copier.Run(b.ctx)
	if err != nil {
		return errors.New(err, fmt.Sprintf("could not copy from %s to %s", srcName, destName))
	}
	return nil
}

func (b *dataLake) List(dirPath string) (dirs []string, objects []string, err error) {

	if dirPath != "" && dirPath[len(dirPath)-1] != '/' {
		dirPath += "/"
	}
	query := &storage.Query{
		Delimiter: "/",
		Prefix:    dirPath,
		Versions:  false,
	}

	iter := b.bucket.Objects(b.ctx, query)
	for {
		objAttrs, iterErr := iter.Next()
		if iterErr == iterator.Done {
			break
		}
		if iterErr != nil {
			err = errors.New(err,"error listing objects from gcs ")
			break
		}
		// If Query.Delimiter is non-empty, some of the ObjectAttrs returned by Next will
		// have a non-empty Prefix field, and a zero value for all other fields. These
		// represent prefixes.
		if objAttrs.Prefix != "" {
			dirs = append(dirs, objAttrs.Prefix)
		} else {
			objects = append(objects, objAttrs.Name)
		}
	}
	if err != nil {
		return nil, nil, errors.New(err, fmt.Sprintf("could not list contents at %s ", dirPath))
	}
	if len(objects) == 0 && len(dirs) == 0 {
		return nil, nil, errors.New(err, fmt.Sprintf("could not find directory %v ", dirPath))
	}
	return
}
