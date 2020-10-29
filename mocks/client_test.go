// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package mocks

import (
	"io"
	"path/filepath"
	"testing"

	"github.com/improbable/imp-billing-datalake/internal/s3"
	"github.com/improbable/imp-billing-datalake/util"
	"github.com/stretchr/testify/assert"
)

func TestDataLakeS3Storage(t *testing.T) {
	bucketName := "imp-billing-datalake-test"
	objectPath := "my-path/my-object"
	objectContent := "just test string"

	FakeS3StorageTestDecorator(bucketName, func(bucketName string, endpoint string) {
		// create s3 datalake
		cfg := util.DataLakeConfig{
			BucketName:         bucketName,
			StorageBackend:     "s3",
			S3CredentialKey:    "test",
			S3CredentialSecret: "test",
			S3Endpoint:         endpoint,
			S3Region:           "test",
		}
		s3.S3ForcePathStyle = true // set s3 ForcePathStyle=true when test
		dl, err := util.NewDataLakeFromConfig(&cfg)
		assert.Nil(t, err)

		// test NewWriter
		wc := dl.NewWriter(objectPath)

		n, err := wc.Write([]byte(objectContent))
		assert.Nil(t, err)
		assert.Equal(t, n, len(objectContent), "write returns len must be equal.")

		err = wc.Close()
		assert.Nil(t, err)

		// test List
		dirs, objs, err := dl.List(objectPath)
		assert.Nil(t, err)
		assert.Equal(t, []string([]string(nil)), dirs)
		assert.Equal(t, []string{objectPath}, objs, "list objs must be equal")

		err = wc.Close()
		assert.Nil(t, err)

		// test Copy
		dstPath := "my-path/my-object2"
		err = dl.Copy(objectPath, dstPath)
		assert.Nil(t, err)

		// test Exists
		exist, err := dl.Exists(dstPath)
		assert.Nil(t, err)
		assert.True(t, exist)

		// test NewReader
		rcr, err := dl.NewReader(dstPath)
		assert.Nil(t, err)

		ret := make([]byte, len(objectContent))
		n, err = rcr.Read(ret)
		assert.Equal(t, io.EOF, err)
		assert.Equal(t, objectContent, string(ret), "read content must be the same.")

		err = rcr.Close()
		assert.Nil(t, err)

		// test Delete
		err = dl.Delete(objectPath)
		assert.Nil(t, err)

		err = dl.Delete(dstPath)
		assert.Nil(t, err)

		exist, err = dl.Exists(dstPath)
		assert.False(t, exist)

	})

}

func TestDataLakeGCSStorage(t *testing.T) {
	bucketName := "sample-bucket"
	objectPath := "my-path/my-object"
	objectContent := "just test strings"

	FakeGCSStorageTestDecorator(bucketName, func(bucketName string, endpoint string) {
		// create gcs datalake
		cfg := util.DataLakeConfig{
			BucketName:             bucketName,
			StorageBackend:         "gcs",
			GCSCredentialJsonBytes: nil,
			GCSEndpoint:            endpoint,
			GCSInsecureSkipVerify:  true,
		}
		dl, err := util.NewDataLakeFromConfig(&cfg)
		assert.Nil(t, err)

		// test NewWriter
		wc := dl.NewWriter(objectPath)
		n, err := wc.Write([]byte(objectContent))
		assert.Nil(t, err)
		assert.Equal(t, n, len(objectContent), "write returns len must be equal.")

		err = wc.Close()
		assert.Nil(t, err)

		// test List
		dirs, objs, err := dl.List(filepath.Dir(objectPath))
		assert.Nil(t, err)
		assert.Equal(t, []string([]string(nil)), dirs)
		assert.True(t, len(objs) > 0)

		err = wc.Close()
		assert.Nil(t, err)

		// test Copy
		dstPath := "my-path/my-object2"
		err = dl.Copy(objectPath, dstPath)
		assert.Nil(t, err)

		// test Exists
		exist, err := dl.Exists(dstPath)
		assert.Nil(t, err)
		assert.True(t, exist)

		// test NewReader
		// Important: because some hard codes make it impossible to access reader TLS content from fake gcs https server,
		// this part has been commented out. And the hard code are from file cloud.google.com/go/storage#NewClient.

		// test Delete
		err = dl.Delete(objectPath)
		assert.Nil(t, err)

		err = dl.Delete(dstPath)
		assert.Nil(t, err)

		exist, err = dl.Exists(dstPath)
		assert.NotNil(t, err)
		assert.False(t, exist)

	})

}
