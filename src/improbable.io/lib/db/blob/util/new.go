// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package blob_util

import (
	"context"

	"improbable.io/lib/db/blob"
	"improbable.io/lib/db/blob/file"
	"improbable.io/lib/db/blob/gcs"
	"improbable.io/lib/db/blob/s3"
	"improbable.io/lib/errors"
	"improbable.io/lib/sharedflags"

	"google.golang.org/grpc/codes"
)

var (
	fBlobBackingStorage = sharedflags.Set.String(
		"backing_blob_storage",
		string(blob.GCS),
		"If you set this option, you can use NewBucketFromFlag to choose your backing storage.\n Available options are defined in `blob.BackingStorage`.\n Keep in mind that each backing storage uses their own set of config flags.")
)

func NewBucketFromFlag(bucketName string) (blob.Bucket, error) {
	return NewBucketFromFlagWithContext(context.Background(), bucketName)
}

func NewBucketFromFlagWithContext(ctx context.Context, bucketName string) (blob.Bucket, error) {
	switch *fBlobBackingStorage {
	case string(blob.GCS):
		return gcs.NewBucket(ctx, bucketName)
	case string(blob.S3):
		return s3.NewBucket(ctx, bucketName)
	case string(blob.Filesystem):
		return file.NewBucket(bucketName)
	default:
		return nil, errors.Newf(nil, codes.InvalidArgument, "Backing storage flag (set to: %s) not set or set incorrectly. Available options are in `blob.BackingStorage`.", *fBlobBackingStorage)
	}
}
