// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package blob

import (
	"context"
	"io"
	"path"

)

// PathPrefixDecorator wraps the supplied bucket so that any paths will be automatically
// proceeded by supplied prefix.
func PathPrefixDecorator(b Bucket, prefix string) Bucket {
	return &pathPrefixDecorator{
		Bucket: b,
		prefix: prefix,
	}
}

type pathPrefixDecorator struct {
	Bucket
	prefix string
}

func (d *pathPrefixDecorator) Exists(p string) (bool, error) {
	return d.Bucket.Exists(path.Join(d.prefix, p))
}

func (d *pathPrefixDecorator) UpdateMeta(p string, attrUpdate ObjectAttributesUpdate) error {
	return d.Bucket.UpdateMeta(path.Join(d.prefix, p), attrUpdate)
}

func (d *pathPrefixDecorator) UpdateACL(p string, acl ACL) error {
	return d.Bucket.UpdateACL(path.Join(d.prefix, p), acl)
}

func (d *pathPrefixDecorator) GetSignedDownloadURL(p string, attr *GetObjectAttributes) (string, error) {
	return d.Bucket.GetSignedDownloadURL(path.Join(d.prefix, p), attr)
}

func (d *pathPrefixDecorator) GetSignedResumableUploadURL(p string, md5 string) (string, error) {
	return d.Bucket.GetSignedResumableUploadURL(path.Join(d.prefix, p), md5)
}

func (d *pathPrefixDecorator) GetSignedCreateMultipartUploadURL(p string) (string, error) {
	return d.Bucket.GetSignedCreateMultipartUploadURL(path.Join(d.prefix, p))
}

func (d *pathPrefixDecorator) GetSignedUploadPartURL(p string, md5 string, uploadId string, partNumber int64) (string, error) {
	return d.Bucket.GetSignedUploadPartURL(path.Join(d.prefix, p), md5, uploadId, partNumber)
}

func (d *pathPrefixDecorator) GetSignedCompleteMultipartUploadURL(p string, uploadId string) (string, error) {
	return d.Bucket.GetSignedCompleteMultipartUploadURL(path.Join(d.prefix, p), uploadId)
}

func (d *pathPrefixDecorator) GetPublicURL(p string) string {
	return d.Bucket.GetPublicURL(path.Join(d.prefix, p))
}

func (d *pathPrefixDecorator) GetSignedUploadURL(p string, md5 string) (string, error) {
	return d.Bucket.GetSignedUploadURL(path.Join(d.prefix, p), md5)
}

func (d *pathPrefixDecorator) Name() string {
	return d.Bucket.Name()
}

func (d *pathPrefixDecorator) Delete(p string) error {
	return d.Bucket.Delete(path.Join(d.prefix, p))
}

func (d *pathPrefixDecorator) NewReader(p string, opts *ReaderOptions) (io.ReadCloser, error) {
	return d.Bucket.NewReader(path.Join(d.prefix, p), nil)
}

func (d *pathPrefixDecorator) NewWriter(ctx context.Context, p string, attrs *ObjectAttributes) io.WriteCloser {
	return d.Bucket.NewWriter(ctx, path.Join(d.prefix, p), attrs)
}

func (d *pathPrefixDecorator) Stat(p string) (*ObjectAttributes, error) {
	return d.Bucket.Stat(path.Join(d.prefix, p))
}

func (d *pathPrefixDecorator) Copy(srcName string, destName string) error {
	return d.Bucket.Copy(path.Join(d.prefix, srcName), path.Join(d.prefix, destName))
}

func (d *pathPrefixDecorator) GetDirContent(ctx context.Context, p string) ([]string, []string, error) {
	return d.Bucket.GetDirContent(ctx, path.Join(d.prefix, p))
}

func (d *pathPrefixDecorator) GetContentWithCommonPrefix(ctx context.Context, prefix string, maxResults int) ([]string, error) {
	return d.Bucket.GetContentWithCommonPrefix(context.TODO(), path.Join(d.prefix, prefix), maxResults)
}
