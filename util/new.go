// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package util

import (
	"context"

	"github.com/max1314/imp-billing-datalake/internal/errors"
	"github.com/max1314/imp-billing-datalake/lib/sharedflags"

	"google.golang.org/grpc/codes"

	blob "github.com/max1314/imp-billing-datalake"
	"github.com/max1314/imp-billing-datalake/gcs"
	"github.com/max1314/imp-billing-datalake/s3"
)

var (
	fBlobDataLakeStorage = sharedflags.Set.String(
		"datalake_blob_storage",
		string(blob.S3),
		"If you set this option, you can use NewBucketFromFlag to choose your backing storage.\n Available options are defined in `blob.BackingStorage`.\n Keep in mind that each backing storage uses their own set of config flags.")
)

func NewDataLakeFromFlag(bucketName string) (blob.DataLake, error) {
	return NewDataLakeFromFlagWithContext(context.Background(), bucketName)
}

func NewDataLakeFromFlagWithContext(ctx context.Context, bucketName string) (blob.DataLake, error) {
	switch *fBlobDataLakeStorage {
	case string(blob.S3):
		return s3.NewDataLake(ctx, bucketName)
	case string(blob.GCS):
		return gcs.NewDataLake(ctx, bucketName)
	default:
		return nil, errors.Newf(nil, codes.InvalidArgument, "Storage Class flag (set to: %s) not set or set incorrectly. Available options are in `blob.BackingStorage`.", *fBlobDataLakeStorage)
	}
}

func SetBlobBackingStorage(s string) {
	fBlobDataLakeStorage = &s
}
