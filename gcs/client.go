// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Contains abstraction over the Google storage API.

package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	blob "github.com/max1314/imp-billing-datalake"
	"github.com/max1314/imp-billing-datalake/internal/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io"
)

// NewDataLake creates a client with explicit values of creds and project
func NewDataLake(ctx context.Context, bucketName string, credsJson []byte) (blob.DataLake, error) {
	bucket := &dataLake{bucketName: bucketName, ctx: ctx}

	conf, err := google.JWTConfigFromJSON(credsJson, storage.ScopeReadWrite, storage.ScopeFullControl)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("storage could not load key: %v. If running locally - have you run git-crypt unlock?", credsJson))
	}
	bucket.urlSigningAccount = conf.Email
	bucket.urlSigningKey = conf.PrivateKey

	bucket.client, err = storage.NewClient(ctx, option.WithTokenSource(conf.TokenSource(ctx)))
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
	//defer reporting.Track(reporting.NewReader, blob.GCS, &err)()

	object := b.bucket.Object(path)
	opts := &blob.ReaderOptions{}
	if opts != nil && opts.GCSReadCompressed {
		object = object.ReadCompressed(true)
	}
	reader, err := object.NewReader(b.ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not create reader for %s ", path))
	}
	return reader, nil
}

func (b *dataLake) NewWriter(path string) io.WriteCloser {
	w := b.bucket.Object(path).NewWriter(b.ctx)
	attrs := &blob.ObjectAttributes{}

	if attrs != nil {
		w.ObjectAttrs.CacheControl = attrs.CacheControl
		w.ObjectAttrs.ContentEncoding = attrs.ContentEncoding
		w.ObjectAttrs.ContentDisposition = attrs.ContentDisposition
		w.ObjectAttrs.ContentLanguage = attrs.ContentLanguage
		w.ObjectAttrs.ContentType = attrs.ContentType
		w.ObjectAttrs.Metadata = attrs.Metadata
	}

	return w
}


func (b *dataLake) Exists(path string) (exists bool, err error) {

	_, err = b.bucket.Object(path).Attrs(b.ctx)
	if err != nil {
		return false, errors.Wrap(err, fmt.Sprintf("could not check if %s exists ", path))
	}
	return true, nil
}

func (b *dataLake) Delete(path string) (err error) {
	err = b.bucket.Object(path).Delete(b.ctx)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("could not delete %s", path))
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
		return errors.Wrap(err, fmt.Sprintf("could not copy from %s to %s", srcName, destName))
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
			err = errors.Wrap(err,"error listing objects from gcs ")
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
		return nil, nil, errors.Wrap(err, fmt.Sprintf("could not list contents at %s ", dirPath))
	}
	if len(objects) == 0 && len(dirs) == 0 {
		return nil, nil, errors.Wrap(err, fmt.Sprintf("could not find directory %v ", dirPath))
	}
	return
}

func toAttributes(gcsAttr *storage.ObjectAttrs) *blob.ObjectAttributes {
	return &blob.ObjectAttributes{
		CacheControl:       gcsAttr.CacheControl,
		ContentEncoding:    gcsAttr.ContentEncoding,
		ContentDisposition: gcsAttr.ContentDisposition,
		ContentLanguage:    gcsAttr.ContentLanguage,
		ContentType:        gcsAttr.ContentType,
		Metadata:           gcsAttr.Metadata,
		MD5:                gcsAttr.MD5,
		Size:               gcsAttr.Size,
		LastModified:       gcsAttr.Updated,
	}
}

