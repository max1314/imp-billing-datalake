// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package imp_billing_datalake

import (
	"io"
)

type BackingStorage string

const (
	GCS BackingStorage = "gcs"
	S3  BackingStorage = "s3"
)

type DataLake interface {
	List(dirPath string) ([]string, []string, error) // return (dirs, objects, err)
	NewReader(path string) (io.ReadCloser, error)
	NewWriter(path string) io.WriteCloser
	Copy(src string, dst string) error
	Delete(path string) error
	Exists(path string) (bool, error)
}
