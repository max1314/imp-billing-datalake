package tests

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/improbable/imp-billing-datalake/util"
)

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