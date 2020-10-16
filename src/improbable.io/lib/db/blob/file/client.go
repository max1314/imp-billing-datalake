// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package file

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"improbable.io/lib/db/blob"
	"improbable.io/lib/db/blob/internal/buffer"
	"improbable.io/lib/errors"
	"improbable.io/lib/util/fileutil"
	"improbable.io/proto/improbable/platform"
)

type bucket struct {
	bucketPath string
}

func NewBucket(name string) (blob.Bucket, error) {
	bucketPath := cleanPath(name)

	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		return nil, errors.Newf(err, errors.NotFound, "the directory %s does not exist", bucketPath)
	}

	return &bucket{bucketPath: bucketPath}, nil
}

func (b *bucket) Name() string {
	return b.bucketPath
}

func (b *bucket) NewReader(path string, opts *blob.ReaderOptions) (io.ReadCloser, error) {
	absolutePath := filepath.Join(b.bucketPath, cleanPath(path))

	file, err := os.Open(absolutePath)
	if os.IsNotExist(err) {
		return nil, errors.Newf(err, errors.NotFound, "the file %s does not exist", absolutePath)
	}
	if err != nil {
		return nil, errors.Newf(err, errors.Internal, "could not open %s for read", absolutePath)
	}

	return file, nil
}

func (b *bucket) Copy(src string, dst string) error {
	fullSrc := filepath.Join(b.bucketPath, cleanPath(src))
	fullDst := filepath.Join(b.bucketPath, cleanPath(dst))

	err := fileutil.CopyFile(fullSrc, fullDst)
	if os.IsNotExist(err) {
		return errors.Newf(err, errors.NotFound, "the file %s does not exist", fullSrc)
	}
	if err != nil {
		return errors.Newf(err, errors.Internal, "failed to copy %s to %s", fullSrc, fullDst)
	}

	return nil
}

func (b *bucket) Delete(path string) error {
	absolutePath := filepath.Join(b.bucketPath, cleanPath(path))
	err := os.Remove(absolutePath)
	if os.IsNotExist(err) {
		return errors.Newf(err, errors.NotFound, "the file %s does not exist", absolutePath)
	}
	if err != nil {
		return errors.Newf(err, errors.Internal, "failed to delete %s", absolutePath)
	}
	return nil
}

func (b *bucket) Exists(path string) (bool, error) {
	absolutePath := filepath.Join(b.bucketPath, cleanPath(path))
	_, err := os.Stat(absolutePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, errors.Newf(err, errors.Internal, "failed to check if file %s exists", absolutePath)
	}
	return true, nil
}

func (b *bucket) NewWriter(ctx context.Context, path string, attrs *blob.ObjectAttributes) io.WriteCloser {
	return buffer.NewClosingBuffer(func(data []byte) error {
		absolutePath := filepath.Join(b.bucketPath, cleanPath(path))
		// make sure directory exists before writing a file into it
		err := os.MkdirAll(filepath.Dir(absolutePath), os.ModePerm)
		if err != nil {
			return errors.Newf(err, errors.Internal, "could not mkdir %s", filepath.Dir(absolutePath))
		}
		err = ioutil.WriteFile(absolutePath, data, os.ModePerm)
		if err != nil {
			return errors.Newf(err, errors.Internal, "could not write to %s", absolutePath)
		}

		return nil
	})
}

func (b *bucket) GetDirContent(ctx context.Context, path string) (dirs []string, objects []string, err error) {
	absolutePath := filepath.Join(b.bucketPath, cleanPath(path))
	dir, err := os.Open(absolutePath)
	if os.IsNotExist(err) {
		return nil, nil, errors.Newf(err, errors.NotFound, "the directory %s does not exist", absolutePath)
	}
	if err != nil {
		return nil, nil, errors.Newf(err, errors.Internal, "could not open %s", absolutePath)
	}

	infos, err := dir.Readdir(0)
	if err != nil {
		return nil, nil, errors.Newf(err, errors.Internal, "could not read files in %s", absolutePath)
	}

	for _, info := range infos {
		// must add path as prefix to be consistent with S3
		if info.IsDir() {
			dirs = append(dirs, path+info.Name())
		} else {
			objects = append(objects, path+info.Name())
		}
	}
	return
}

func (b *bucket) GetContentWithCommonPrefix(ctx context.Context, prefix string, maxResults int) ([]string, error) {
	absolutePath := filepath.Join(b.bucketPath, cleanPath(filepath.Dir(prefix)))
	entries := []string{}
	err := filepath.Walk(absolutePath, func(path string, info os.FileInfo, err error) error {
		if maxResults > 0 && len(entries) >= maxResults {
			// we already gathered required number, skip everything afterwards
			return filepath.SkipDir
		}
		if os.IsNotExist(err) {
			// file just got deleted before this func run, ignore it
			return nil
		}
		if err != nil {
			return errors.Newf(err, errors.Internal, "could not read %s", absolutePath)
		}
		relPath, err := filepath.Rel(b.bucketPath, path)
		if err != nil {
			return errors.Newf(err, errors.Internal, "internal reading problem at %s", absolutePath)
		}
		if info.IsDir() {
			relPath += "/"
		}
		if strings.HasPrefix(relPath, prefix) {
			entries = append(entries, relPath)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Newf(err, errors.Internal, "failed to traver the folder at %s", absolutePath)
	}
	return entries, nil
}

func (b *bucket) UpdateMeta(path string, attrUpdate blob.ObjectAttributesUpdate) error {
	return errors.New(nil, errors.Unimplemented, "this method is not supported by a local filesystem blob client")
}

func (b *bucket) Stat(path string) (*blob.ObjectAttributes, error) {
	return nil, errors.New(nil, errors.Unimplemented, "this method is not supported by a local filesystem blob client")
}

func (b *bucket) UpdateACL(path string, acl blob.ACL) error {
	return errors.New(nil, errors.Unimplemented, "this method is not supported by a local filesystem blob client")
}

func (b *bucket) GetSignedDownloadURL(path string, _ *blob.GetObjectAttributes) (string, error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by a local filesystem blob client")
}

func (b *bucket) GetSupportedFileStorageProvider() improbable_platform.StorageProvider {
	return improbable_platform.StorageProvider_UNKNOWN
}

func (b *bucket) GetSignedUploadURL(path string, md5 string) (string, error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by a local filesystem blob client")
}

func (b *bucket) GetSignedResumableUploadURL(path string, md5 string) (string, error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by a local filesystem blob client")
}

func (b *bucket) GetSignedCreateMultipartUploadURL(path string) (string, error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by a local filesystem blob client")
}

func (b *bucket) GetSignedUploadPartURL(path string, md5 string, uploadId string, partNumber int64) (string, error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by a local filesystem blob client")
}

func (b *bucket) GetSignedCompleteMultipartUploadURL(path string, uploadId string) (string, error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by a local filesystem blob client")
}

func (b *bucket) GetPublicURL(path string) string {
	return ""
}

func cleanPath(path string) string {
	if runtime.GOOS == "windows" {
		return filepath.FromSlash(path)
	}

	return filepath.ToSlash(path)
}
