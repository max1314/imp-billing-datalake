// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package blob

import (
	"context"
	"crypto/md5"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"

	"improbable.io/lib/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestBucket() Bucket {
	bucket, _ := NewFakeBucket("bucketName", map[string]*FakeBlob{
		"some_file": {Contents: "some_file_contents"},
	})
	return bucket
}

func TestFakeClient_FileExists(t *testing.T) {
	bucket := CreateTestBucket()
	present, err := bucket.Exists("some_file")

	assert.True(t, err == nil, "no error")
	assert.True(t, present, "file exists")
}

func TestFakeClient_FileDoesNotExist(t *testing.T) {
	bucket := CreateTestBucket()
	present, err := bucket.Exists("another_file")

	assert.True(t, err == nil, "no error")
	assert.False(t, present, "file does not exist")
}

func TestFakeClient_NewReader(t *testing.T) {
	bucket := CreateTestBucket()

	reader, err := bucket.NewReader("some_file", nil)
	if reader != nil {
		defer reader.Close()
	}
	assert.True(t, err == nil, "no error on NewReader")

	actual, err := ioutil.ReadAll(reader)
	assert.True(t, err == nil, "no error on ReadAll")
	assert.Equal(t, string(actual), "some_file_contents", "reader returned")
}

func TestFakeClient_Delete(t *testing.T) {
	bucket := CreateTestBucket()

	bucket.Delete("some_file")

	present, err := bucket.Exists("some_file")

	assert.True(t, err == nil, "no error")
	assert.False(t, present, "file deleted")
}

func TestFakeClient_NewWriter(t *testing.T) {
	bucket := CreateTestBucket()

	w := bucket.NewWriter(context.TODO(), "new_file", nil)
	io.WriteString(w, "expected_content")
	w.Close()

	present, err := bucket.Exists("new_file")

	assert.True(t, err == nil, "no error")
	assert.True(t, present, "file created")

	reader, err := bucket.NewReader("new_file", nil)
	if reader != nil {
		defer reader.Close()
	}
	assert.True(t, err == nil, "no error on NewReader")

	actual, err := ioutil.ReadAll(reader)
	assert.True(t, err == nil, "no error on ReadAll")
	assert.Equal(t, string(actual), "expected_content", "reader returned")
}

func TestFakeClient_NewWriter_SetsAttributes(t *testing.T) {
	bucket := CreateTestBucket()
	attrs := &ObjectAttributes{
		ContentType:        "text/json",
		ContentEncoding:    "gzip",
		ContentDisposition: `attachment; filename="filename.jpg"`,
		CacheControl:       "max-age=360",
		ContentLanguage:    "english",
		Metadata: map[string]string{
			"foo": "bar",
		},
	}
	w := bucket.NewWriter(context.TODO(), "new_file", attrs)
	_, err := io.WriteString(w, "test")
	require.NoError(t, err, "must not fail to write")

	require.NoError(t, w.Close(), "must not fail to close")

	fileAttrs, err := bucket.Stat("new_file")
	require.NoError(t, err)

	require.Equal(t, attrs.Metadata, fileAttrs.Metadata, "metadata must match")
	require.Equal(t, attrs.ContentType, fileAttrs.ContentType, "ContentType must match")
	require.Equal(t, attrs.ContentEncoding, fileAttrs.ContentEncoding, "ContentEncoding must match")
	require.Equal(t, attrs.ContentDisposition, fileAttrs.ContentDisposition, "ContentDisposition must match")
	require.Equal(t, attrs.CacheControl, fileAttrs.CacheControl, "CacheControl must match")
	require.Equal(t, attrs.ContentLanguage, fileAttrs.ContentLanguage, "ContentLanguage must match")
}

func TestFakeClient_WriteAndUpdateMeta(t *testing.T) {
	bucket := CreateTestBucket()
	path := "new_file"

	w := bucket.NewWriter(context.TODO(), path, nil)
	_, err := io.WriteString(w, "expected_content")
	require.NoError(t, err, "must not fail to write")
	require.NoError(t, w.Close(), "must not fail to close")

	present, err := bucket.Exists(path)
	require.NoError(t, err, "must not fail to check existence of file")
	assert.True(t, present, "file must exist")

	err = bucket.UpdateMeta(path, ObjectAttributesUpdate{})
	require.NoError(t, err, "must not fail to update attributes of the file previously written")
}

func TestFakeClient_StatsErrors_WhenObjectMissing(t *testing.T) {
	bucket := CreateTestBucket()

	_, err := bucket.Stat("some_path")
	require.Error(t, err, "must fail because object does not exist")
}

func TestFakeClient_StatsReturnsDefaultAttributes_WhenNoAttributesSet(t *testing.T) {
	bucket := CreateTestBucket()

	w := bucket.NewWriter(context.TODO(), "some_path", nil)
	_, err := w.Write([]byte("some_data"))
	require.NoError(t, err)
	require.NoError(t, w.Close())

	attrs, err := bucket.Stat("some_path")
	require.NoError(t, err, "must not fail when attributes not set but object exists")

	hash := md5.Sum([]byte("some_data"))
	assert.Equal(t, attrs, &ObjectAttributes{
		MD5: hash[:],
	})
}

func TestFakeClient_GetDirContent_ReturnsSubDirsAndObjects_WhenProvidedDirPath(t *testing.T) {
	bucket, err := NewFakeBucket("/", map[string]*FakeBlob{
		filepath.Join("root", "foo.json"):        {},
		filepath.Join("root", "bar", "baz.json"): {},
	})
	require.NoError(t, err)
	dirs, objs, err := bucket.GetDirContent(context.TODO(), "root")
	require.NoError(t, err)
	require.Equal(t, []string{filepath.Join("root", "bar")}, dirs)
	require.Equal(t, []string{filepath.Join("root", "foo.json")}, objs)
}

func TestFakeClient_GetDirContent_ReturnInvalidArgumentError_WhenProvidedObjectPath(t *testing.T) {
	bucket, err := NewFakeBucket("/", map[string]*FakeBlob{
		"root/foo.json":     {},
		"root/bar/baz.json": {},
	})
	require.NoError(t, err)
	dirs, objs, err := bucket.GetDirContent(context.TODO(), "root/foo.json")
	require.Nil(t, dirs)
	require.Nil(t, objs)
	require.Equal(t, errors.InvalidArgument, errors.Code(err))
}

func TestFakeClient_GetDirContent_ReturnsNotFoundError_WhenProvidedUnknownPath(t *testing.T) {
	bucket, err := NewFakeBucket("/", map[string]*FakeBlob{
		"root/foo.json":     {},
		"root/bar/baz.json": {},
	})
	require.NoError(t, err)
	dirs, objs, err := bucket.GetDirContent(context.TODO(), "user")
	require.Nil(t, dirs)
	require.Nil(t, objs)
	require.Equal(t, errors.NotFound, errors.Code(err))
}
