// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Provides a Fake Bucket implementation. It provides an in-memory blob storage
// suitable for use as a test-double.
package blob

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"improbable.io/lib/db/blob/internal/buffer"
	"improbable.io/lib/errors"

	"google.golang.org/grpc/codes"
)

// NewFakeBucket returns a new Bucket instance which contains the supplied map of
// objects.
func NewFakeBucket(path string, initialObjects map[string]*FakeBlob) (*FakeBucket, error) {
	return &FakeBucket{
		bucketName:        path,
		objects:           initialObjects,
		signedDownloadURl: FakeSignedDownloadURL,
	}, nil
}

type FakeBlob struct {
	// Contents are the contents of the blob
	Contents string
	// Attrs represent the object attributes of blob
	Attrs *ObjectAttributes
	// ACLs represent the access level of the object
	Acl ACL
}

type FakeBucket struct {
	sync.RWMutex
	bucketName        string
	objects           map[string]*FakeBlob
	signedDownloadURl string
}

func (b *FakeBucket) Exists(path string) (bool, error) {
	b.RLock()
	defer b.RUnlock()
	_, ok := b.objects[path]
	return ok, nil
}

func (b *FakeBucket) Stat(path string) (*ObjectAttributes, error) {
	b.RLock()
	defer b.RUnlock()
	obj, ok := b.objects[path]
	if !ok {
		return nil, errors.New(nil, codes.NotFound, "object does not exist")
	}

	if obj.Attrs == nil {
		return &ObjectAttributes{}, nil
	}

	return obj.Attrs, nil
}

func (b *FakeBucket) UpdateMeta(path string, attrUpdate ObjectAttributesUpdate) error {
	b.RLock()
	defer b.RUnlock()

	obj, ok := b.objects[path]
	if !ok {
		return errors.New(nil, errors.NotFound, "Object not found")
	}

	if attrUpdate.CacheControl != nil {
		obj.Attrs.CacheControl = *attrUpdate.CacheControl
	}
	if attrUpdate.ContentEncoding != nil {
		obj.Attrs.ContentEncoding = *attrUpdate.ContentEncoding
	}
	if attrUpdate.ContentDisposition != nil {
		obj.Attrs.ContentDisposition = *attrUpdate.ContentDisposition
	}
	if attrUpdate.ContentLanguage != nil {
		obj.Attrs.ContentLanguage = *attrUpdate.ContentLanguage
	}
	if attrUpdate.ContentType != nil {
		obj.Attrs.ContentType = *attrUpdate.ContentType
	}
	if attrUpdate.Metadata != nil {
		obj.Attrs.Metadata = attrUpdate.Metadata
	}

	return nil
}

func (b *FakeBucket) UpdateACL(path string, acl ACL) error {
	b.RLock()
	defer b.RUnlock()

	obj, ok := b.objects[path]
	if !ok {
		return errors.Newf(nil, errors.NotFound, "not found: %s/%s", b.bucketName, path)
	}
	obj.Acl = acl
	return nil
}

func (b *FakeBucket) NewReader(path string, opts *ReaderOptions) (io.ReadCloser, error) {
	b.RLock()
	defer b.RUnlock()
	obj, ok := b.objects[path]
	if !ok {
		return nil, errors.Newf(nil, errors.NotFound, "not found: %s/%s", b.bucketName, path)
	}
	if obj.Acl == AUTHENTICATED_READ {
		return nil, errors.Newf(nil, codes.Unimplemented, "authentication check is not supported in fake client")
	}
	if obj.Acl == PRIVATE {
		return nil, errors.Newf(nil, errors.PermissionDenied, "item is private and cannot be fetched")
	}
	reader := ioutil.NopCloser(bytes.NewBufferString(obj.Contents))
	return reader, nil
}

func (b *FakeBucket) NewWriter(ctx context.Context, path string, attrs *ObjectAttributes) io.WriteCloser {
	return buffer.NewClosingBuffer(func(data []byte) error {
		b.Lock()
		hash := md5.Sum(data)

		objAttrs := attrs
		if objAttrs == nil {
			objAttrs = &ObjectAttributes{}
		}
		objAttrs.MD5 = hash[:]

		b.objects[path] = &FakeBlob{
			Contents: string(data[:]),
			Attrs:    objAttrs,
			Acl:      UNKNOWN,
		}
		b.Unlock()
		return nil
	})
}

func (b *FakeBucket) Delete(path string) error {
	b.Lock()
	defer b.Unlock()
	if _, ok := b.objects[path]; !ok {
		return errors.Newf(nil, errors.NotFound, "Item %s not in bucket!", path)
	}
	delete(b.objects, path)
	return nil
}

func (b *FakeBucket) Copy(src string, dst string) error {
	return nil
}

func (b *FakeBucket) Name() string {
	return b.bucketName
}

func (b *FakeBucket) GetSignedUploadURL(path string, md5 string) (string, error) {
	return fmt.Sprintf("https://fake.upload.improbable.io/%s/%s", b.bucketName, path), nil
}

func (b *FakeBucket) GetSignedResumableUploadURL(path string, md5 string) (string, error) {
	return fmt.Sprintf("https://fake.upload.improbable.io/%s/%s", b.bucketName, path), nil
}

func (b *FakeBucket) GetSignedCreateMultipartUploadURL(path string) (url string, err error) {
	return fmt.Sprintf("https://fake.upload.improbable.io/%s/%s", b.bucketName, path), nil
}

func (b *FakeBucket) GetSignedUploadPartURL(path string, md5 string, uploadId string, partNumber int64) (url string, err error) {
	return fmt.Sprintf("https://fake.upload.improbable.io/%s/%s", b.bucketName, path), nil
}

func (b *FakeBucket) GetSignedCompleteMultipartUploadURL(path string, uploadId string) (url string, err error) {
	return fmt.Sprintf("https://fake.upload.improbable.io/%s/%s", b.bucketName, path), nil
}

const FakeSignedDownloadURL = "https://fake.download.improbable.io"

func (b *FakeBucket) SetSignedDownloadURL(url string) {
	b.signedDownloadURl = url
}

func (b *FakeBucket) GetSignedDownloadURL(path string, attr *GetObjectAttributes) (string, error) {
	return fmt.Sprintf("%s/%s/%s", b.signedDownloadURl, b.bucketName, path), nil
}

const FakeUnsignedURL = "https://fake.public.improbable.io"

func (b *FakeBucket) GetPublicURL(path string) string {
	return fmt.Sprintf("%s/%s/%s", FakeUnsignedURL, b.bucketName, path)
}

func (b *FakeBucket) GetDirContent(ctx context.Context, dirPath string) (dirs []string, objects []string, err error) {
	for objPath := range b.objects {
		if strings.HasPrefix(objPath, dirPath) {
			relPath := strings.TrimPrefix(strings.TrimPrefix(objPath, dirPath), "/")
			parts := strings.Split(relPath, "/")
			if parts[0] == "" {
				// current object sits at requested path, so requested path was not a directory
				return nil, nil, errors.Newf(nil, errors.InvalidArgument, "%s is not a directory", dirPath)
			} else if len(parts) == 1 && parts[0] != "" {
				// current object is one level deeper than requested path, so collect the object
				// we must collect objPath to keep consistency with S3 implementation
				objects = append(objects, objPath)
			} else if len(parts) > 1 {
				// current object is more than one level deeper than requested path, so collect the intermediate dir
				dirs = append(dirs, filepath.Join(dirPath, parts[0]))
			}
		}
	}
	if len(objects) == 0 && len(dirs) == 0 {
		return nil, nil, errors.Newf(nil, errors.NotFound, "no content found at %s", dirPath)
	}
	return dirs, objects, nil
}

func (b *FakeBucket) GetContentWithCommonPrefix(ctx context.Context, prefix string, maxResults int) ([]string, error) {
	var keys []string
	for k := range b.objects {
		if strings.HasPrefix(k, prefix) {
			keys = append(keys, k)
		}
	}

	if maxResults != 0 && len(keys) > maxResults {
		keys = keys[:maxResults]
	}

	return keys, nil
}
