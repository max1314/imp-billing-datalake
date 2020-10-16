// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package gcs

import (
	"context"
	"io"

	"improbable.io/lib/db/blob"
	"improbable.io/lib/db/blob/gcs"
)

type dataLake struct {
	bucket blob.Bucket
}

func NewDataLake(ctx context.Context, bucketName string) (*DataLake, error) {
	b, err := gcs.NewBucket(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	return &dataLake{bucket: b}, nil
}

func (dl *dataLake) List(dirPath string) (dirs []string, objects []string, err error) {
	ctx := context.Background()
	return dl.bucket.GetDirContent(ctx, dirPath)
}

func (dl *dataLake) NewReader(path string) (io.ReadCloser, error) {
	opts := &datalake.ReaderOptions{}
	return dl.bucket.NewReader(path, opts)
}

func (dl *dataLake) NewWriter(path string) io.WriteCloser {
	ctx := context.Background()
	attrs := &datalake.ObjectAttributes{}
	return dl.bucket.NewWriter(ctx, path, attrs)
}

func (dl *dataLake) Copy(src string, dst string) error {
	return dl.bucket.Copy(src, dst)
}

func (dl *dataLake) Delete(path string) error {
	return dl.bucket.Delete(path)
}

func (dl *dataLake) Exists(path string) (bool, error) {
	return dl.bucket.Exists(path)
}
