// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package imp_billing_datalake

import (
	"io"

	"improbable.io/lib/db/blob"
)

type StorageClass = blob.BackingStorage

const (
	GCS StorageClass = "gcs"
	S3  StorageClass = "s3"
)

// reuse the lib/blob structs
type ObjectAttributes = blob.ObjectAttributes
type ReaderOptions = blob.ReaderOptions

type DataLake interface {
	List(dirPath string) ([]string, []string, error) // return (dirs, objects, err)
	NewReader(path string) (io.ReadCloser, error)
	NewWriter(path string) io.WriteCloser
	Copy(src string, dst string) error
	Delete(path string) error
	Exists(path string) (bool, error)
}
