// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package blob

import (
	"context"
	"io"
	"time"
)

type ACL int

const (
	UNKNOWN ACL = iota
	PRIVATE
	AUTHENTICATED_READ
	PUBLIC_READ
)

type BackingStorage string

const (
	GCS        BackingStorage = "gcs"
	S3         BackingStorage = "s3"
	Filesystem BackingStorage = "filesystem"
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

type BucketReader interface {
	NewReader(path string, opts *ReaderOptions) (io.ReadCloser, error)
	GetPublicURL(path string) string
}

// Provides an interface abstraction over upstream Storage API that acts on a specific bucket.
type Bucket interface {
	BucketReader

	Copy(src string, dst string) error
	Delete(path string) error
	Exists(path string) (bool, error)
	GetSignedDownloadURL(path string, attr *GetObjectAttributes) (string, error)
	// GCS and S3
	GetSignedUploadURL(path string, md5 string) (string, error)
	// GCS only
	GetSignedResumableUploadURL(path string, md5 string) (string, error)
	// S3 only
	// Generate a presigned CreateMultipartUpload request url (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)
	GetSignedCreateMultipartUploadURL(path string) (url string, err error)
	// S3 only
	// Generate a presigned UploadPart request url (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)
	GetSignedUploadPartURL(path string, md5 string, uploadId string, partNumber int64) (url string, err error)
	// S3 only
	// Generate a presigned CompleteMultipartUpload request url (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)
	GetSignedCompleteMultipartUploadURL(path string, uploadId string) (url string, err error)
	Name() string
	NewWriter(ctx context.Context, path string, attrs *ObjectAttributes) io.WriteCloser
	Stat(path string) (*ObjectAttributes, error)
	UpdateMeta(path string, attrUpdate ObjectAttributesUpdate) error
	UpdateACL(path string, acl ACL) error

	// Get the "content of a directory" given a common path prefix as a list directory names
	// and a list of object names found. It throws errors if the nothing matches the directory
	// prefix or if the bucket cannot be accessed.
	GetDirContent(ctx context.Context, path string) (dirs []string, objects []string, err error)
	// Return all objects with a common prefix. Notice this will list all files inside of sub
	// "directories". If maxResults is 0, it'll return all results.
	GetContentWithCommonPrefix(ctx context.Context, prefix string, maxResults int) ([]string, error)
}
