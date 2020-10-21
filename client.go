// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package imp_billing_datalake

import (
	"io"
	"time"
)

type BackingStorage string

const (
	GCS BackingStorage = "gcs"
	S3  BackingStorage = "s3"
)

type ObjectAttributesUpdate struct {
	CacheControl       *string
	ContentEncoding    *string
	ContentDisposition *string
	ContentLanguage    *string
	ContentType        *string
	Metadata           map[string]string
}

type ObjectAttributes struct {
	CacheControl       string
	ContentEncoding    string
	ContentDisposition string
	ContentLanguage    string
	ContentType        string
	Metadata           map[string]string
	MD5                []byte
	Size               int64
	LastModified       time.Time
}

type GetObjectAttributes struct {
	// ResponseContentDisposition is only used in S3
	ResponseContentDisposition string
}

type ReaderOptions struct {
	GCSReadCompressed bool
}


type DataLake interface {
	List(dirPath string) ([]string, []string, error) // return (dirs, objects, err)
	NewReader(path string) (io.ReadCloser, error)
	NewWriter(path string) io.WriteCloser
	Copy(src string, dst string) error
	Delete(path string) error
	Exists(path string) (bool, error)
}
