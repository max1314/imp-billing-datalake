// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package util

import (
	"context"

	"improbable.io/lib/errors"
	"improbable.io/lib/sharedflags"

	"google.golang.org/grpc/codes"

	"imp-billing-datalake/gcs"
	"imp-billing-datalake/s3"
	datalake "imp-billing-datalake"
)

var (
	fBlobDataLakeStorage = sharedflags.Set.String(
		"datalake_blob_storage",
		string(datalake.S3),
		"If you set this option, you can use NewBucketFromFlag to choose your backing storage.\n Available options are defined in `blob.BackingStorage`.\n Keep in mind that each backing storage uses their own set of config flags.")
)

func NewDataLakeFromFlag(bucketName string) (DataLake, error) {
	return NewDataLakeFromFlagWithContext(context.Background(), bucketName)
}

func NewDataLakeFromFlagWithContext(ctx context.Context, bucketName string) (*datalake.DataLake, error) {
	switch *fBlobDataLakeStorage {
	case string(datalake.S3):
		return s3.NewDataLake(ctx, bucketName)
	case string(datalake.GCS):
		return gcs.NewDataLake(ctx, bucketName)
	default:
		return nil, errors.Newf(nil, codes.InvalidArgument, "Storage Class flag (set to: %s) not set or set incorrectly. Available options are in `blob.BackingStorage`.", *fBlobDataLakeStorage)
	}
}

func SetBlobBackingStorage(s string) {
	fBlobDataLakeStorage = &s
}