// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Contains abstraction over the Google storage API.

package gcs

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"improbable.io/lib/db/blob"
	"improbable.io/lib/db/blob/internal/reporting"
	"improbable.io/lib/errors"
	"improbable.io/lib/sharedflags"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	fSignedURLExpirationTime = sharedflags.Set.Duration(
		"gcs_signed_url_expiration_time",
		20*time.Minute,
		"The expiration time for signed URLs for the storage API in minutes")

	fCredsType = sharedflags.Set.String(
		"gcs_credentials_type",
		"json",
		"The type of credentials for Cloud Storage: [json, implicit, auto]. "+
			"'json' instructs reading service account from --gcs_credentials_secret, "+
			"'implicit' will use GCE token source and "+
			"'auto' attempts to use gcloud application default credentials. "+
			"URL signing options are not available when using implicit credentials.")

	fCredsSecret = sharedflags.Set.String(
		"gcs_credentials_secret",
		os.ExpandEnv("$GOPATH/keys/golden-bucket-91717_readonly_signing_account.json"),
		"Path to secrets when type is json, with env expansion for Cloud Storage.")

	fCredsAccount = sharedflags.Set.String(
		"gcs_credentials_account",
		"default", // GCE default compute account.
		"Service Account to used for accessing Cloud Storage, only needed when type is pem or implicit (GCE).")
)

// NewBucket returns a new bucket client built based off flagz for the given bucket.
// You probably want to use improbable.io/lib/db/blob/util.NewBucketFromFlag instead for easier support for multiple clouds.
func NewBucket(ctx context.Context, bucketName string) (blob.Bucket, error) {
	return NewExplicitBucket(ctx, bucketName, *fCredsType, *fCredsSecret, *fCredsAccount)
}

// NewExplicitBucket creates a client with explicit values of creds and project
func NewExplicitBucket(ctx context.Context, bucketName string, credsType string, credsSecretPath string, credsAccount string) (blob.Bucket, error) {
	bucket := &bucket{bucketName: bucketName, ctx: ctx}
	switch credsType {
	case "json":
		credsJson, err := ioutil.ReadFile(credsSecretPath)
		if err != nil {
			return nil, errors.Newf(err, errors.Internal, "could not read private key file: %v. If running locally - have you run git-crypt unlock?", err)
		}
		conf, err := google.JWTConfigFromJSON(credsJson, storage.ScopeReadWrite, storage.ScopeFullControl)
		if err != nil {
			return nil, errors.Newf(err, errors.Internal, "storage could not load key: %v. If running locally - have you run git-crypt unlock?", err)
		}
		bucket.urlSigningAccount = conf.Email
		bucket.urlSigningKey = conf.PrivateKey

		bucket.client, err = storage.NewClient(ctx, option.WithTokenSource(conf.TokenSource(ctx)))
		if err != nil {
			return nil, err
		}
	case "implicit":
		source := google.ComputeTokenSource(credsAccount)

		var err error
		bucket.client, err = storage.NewClient(ctx, option.WithTokenSource(source))
		if err != nil {
			return nil, err
		}
	case "auto":
		var err error
		bucket.client, err = storage.NewClient(ctx)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(nil, errors.Internal, "currently only support json credentials")
	}

	bucket.bucket = bucket.client.Bucket(bucketName)
	return bucket, nil
}

// NewBucketWithTokenSource creates a new GCS bucket from OIDC token source.
func NewBucketWithTokenSource(ctx context.Context, bucketName string, source oauth2.TokenSource) (blob.Bucket, error) {
	client, err := storage.NewClient(ctx, option.WithTokenSource(source))
	if err != nil {
		return nil, err
	}

	return &bucket{
		bucket:     client.Bucket(bucketName),
		ctx:        ctx,
		client:     client,
		bucketName: bucketName,
	}, nil
}

type bucket struct {
	ctx               context.Context
	client            *storage.Client
	bucket            *storage.BucketHandle
	bucketName        string
	urlSigningAccount string
	urlSigningKey     []byte
}

func (b *bucket) NewReader(path string, opts *blob.ReaderOptions) (r io.ReadCloser, err error) {
	defer reporting.Track(reporting.NewReader, blob.GCS, &err)()

	object := b.bucket.Object(path)
	if opts != nil && opts.GCSReadCompressed {
		object = object.ReadCompressed(true)
	}
	reader, err := object.NewReader(b.ctx)
	if err != nil {
		return nil, errors.Wrapf(toStandardError(err), "Could not create reader for %s", path)
	}
	return reader, nil
}

func (b *bucket) NewWriter(ctx context.Context, path string, attrs *blob.ObjectAttributes) io.WriteCloser {
	w := b.bucket.Object(path).NewWriter(ctx)

	if attrs != nil {
		w.ObjectAttrs.CacheControl = attrs.CacheControl
		w.ObjectAttrs.ContentEncoding = attrs.ContentEncoding
		w.ObjectAttrs.ContentDisposition = attrs.ContentDisposition
		w.ObjectAttrs.ContentLanguage = attrs.ContentLanguage
		w.ObjectAttrs.ContentType = attrs.ContentType
		w.ObjectAttrs.Metadata = attrs.Metadata
	}

	return w
}

func (b *bucket) Stat(path string) (attrs *blob.ObjectAttributes, err error) {
	defer reporting.Track(reporting.Stat, blob.GCS, &err)()
	obj, err := b.bucket.Object(path).Attrs(b.ctx)
	if err != nil {
		return nil, errors.Wrapf(toStandardError(err), "could not get stat for %s", path)
	}
	return toAttributes(obj), nil
}

func (b *bucket) UpdateMeta(path string, attrUpdate blob.ObjectAttributesUpdate) (err error) {
	defer reporting.Track(reporting.UpdateMeta, blob.GCS, &err)()

	storageAttr := storage.ObjectAttrsToUpdate{}
	if attrUpdate.CacheControl != nil {
		storageAttr.CacheControl = *attrUpdate.CacheControl
	}
	if attrUpdate.ContentEncoding != nil {
		storageAttr.ContentEncoding = *attrUpdate.ContentEncoding
	}
	if attrUpdate.ContentDisposition != nil {
		storageAttr.ContentDisposition = *attrUpdate.ContentDisposition
	}
	if attrUpdate.ContentLanguage != nil {
		storageAttr.ContentLanguage = *attrUpdate.ContentLanguage
	}
	if attrUpdate.ContentType == nil || *attrUpdate.ContentType == "" {
		// GCS client libraries will start to fail in an infinite loop if content type is not set.
		return errors.Newf(nil, errors.Internal, "ContentType for %q must be set", path)
	}
	storageAttr.ContentType = *attrUpdate.ContentType
	if attrUpdate.Metadata != nil {
		storageAttr.Metadata = attrUpdate.Metadata
	}

	_, err = b.bucket.Object(path).Update(b.ctx, storageAttr)
	if err != nil {
		return errors.Wrap(toStandardError(err), "cloud not update meta for %s")
	}
	return nil
}

func (b *bucket) UpdateACL(path string, acl blob.ACL) (err error) {
	defer reporting.Track(reporting.UpdateACL, blob.GCS, &err)()

	_, err = b.bucket.Object(path).Update(b.ctx, storage.ObjectAttrsToUpdate{
		PredefinedACL: aclToGcsRules(acl),
	})
	if err != nil {
		return errors.Wrapf(toStandardError(err), "could not update acl for %s", path)
	}
	return nil
}

func (b *bucket) Exists(path string) (exists bool, err error) {
	defer reporting.Track(reporting.Exists, blob.GCS, &err)()

	_, err = b.bucket.Object(path).Attrs(b.ctx)
	if err != nil {
		if errors.Categorise(toStandardError(err)) != errors.System {
			return false, nil
		}
		return false, errors.Wrapf(toStandardError(err), "could not check if %s exists", path)
	}
	return true, nil
}

func (b *bucket) GetPublicURL(path string) string {
	return fmt.Sprintf("https://commondatastorage.googleapis.com/%v/%v", b.bucketName, path)
}

func (b *bucket) GetSignedDownloadURL(path string, _ *blob.GetObjectAttributes) (url string, err error) {
	defer reporting.Track(reporting.GetSignedDownloadURL, blob.GCS, &err)()
	opt, err := b.signedURLOptFromJwtConfig("GET", time.Now().Add(*fSignedURLExpirationTime))
	if err != nil {
		return "", errors.Wrapf(toStandardError(err), "could not sign a download url for %s", path)
	}
	signedURL, err := storage.SignedURL(b.bucketName, path, opt)
	if err != nil {
		return "", errors.Wrapf(toStandardError(err), "could not sign a download url for %s", path)
	}
	return signedURL, nil
}

// Generate a signed URL using the provided MD5 to allow an upload via PUT.
// NOTE: The MD5 must be the Base64 encoded string of the MD5 digest byte array.
func (b *bucket) GetSignedUploadURL(path string, md5 string) (url string, err error) {
	defer reporting.Track(reporting.GetSignedUploadURL, blob.GCS, &err)()
	opts, err := b.signedURLOptFromJwtConfig("PUT", time.Now().Add(*fSignedURLExpirationTime))
	if err != nil {
		return "", errors.Wrapf(toStandardError(err), "Could not sign a upload url for %s", path)
	}

	if md5 == "" {
		return "", errors.New(nil, errors.InvalidArgument, "a base64 encoded MD5 of the file to be uploaded must be included")
	}

	opts.MD5 = md5
	signedURL, err := storage.SignedURL(b.bucketName, path, opts)
	if err != nil {
		return "", errors.Wrapf(toStandardError(err), "Could not sign a upload url for %s", path)
	}
	return signedURL, nil
}

// Generate a signed URL using the MD5 provided to initiate a resumable upload via a POST request.
// According to GCS's XML API (see https://cloud.google.com/storage/docs/xml-api/resumable-upload).
// NOTE: The MD5 must be the Base64 encoded string of the MD5 digest byte array.
func (b *bucket) GetSignedResumableUploadURL(path string, md5 string) (url string, err error) {
	defer reporting.Track(reporting.GetSignedResumableUploadURL, blob.GCS, &err)()
	opts, err := b.signedURLOptFromJwtConfig("POST", time.Now().Add(*fSignedURLExpirationTime))
	if err != nil {
		return "", errors.Wrapf(toStandardError(err), "could not sign a resumable upload url for %s", path)
	}
	if md5 == "" {
		return "", errors.New(nil, errors.InvalidArgument, "a base64 encoded MD5 of the file to be uploaded must be included")
	}
	opts.MD5 = md5
	opts.Headers = append(opts.Headers, "x-goog-resumable:start")
	signedURL, err := storage.SignedURL(b.bucketName, path, opts)
	if err != nil {
		return "", errors.Wrapf(toStandardError(err), "could not sign a resumable upload url for %s", path)
	}
	return signedURL, nil
}

func (b *bucket) signedURLOptFromJwtConfig(method string, expires time.Time) (*storage.SignedURLOptions, error) {
	if b.urlSigningAccount == "" || b.urlSigningKey == nil || len(b.urlSigningKey) == 0 {
		return nil, errors.New(nil, errors.Internal, "no url signing key or account was provided for the bucket")
	}
	return &storage.SignedURLOptions{
		GoogleAccessID: b.urlSigningAccount,
		PrivateKey:     b.urlSigningKey,
		Method:         method,
		Expires:        expires,
	}, nil
}

func (b *bucket) GetSignedCreateMultipartUploadURL(path string) (url string, err error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by this blob client type")
}

func (b *bucket) GetSignedUploadPartURL(path string, md5 string, uploadId string, partNumber int64) (url string, err error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by this blob client type")
}

func (b *bucket) GetSignedCompleteMultipartUploadURL(path string, uploadId string) (url string, err error) {
	return "", errors.New(nil, errors.Unimplemented, "this method is not supported by this blob client type")
}

func (b *bucket) Delete(path string) (err error) {
	defer reporting.Track(reporting.Delete, blob.GCS, &err)()
	err = b.bucket.Object(path).Delete(b.ctx)
	if err != nil {
		return errors.Wrapf(toStandardError(err), "could not delete %s", path)
	}
	return nil
}

// Copy calls google api to do a copy
func (b *bucket) Copy(srcName string, destName string) (err error) {
	defer reporting.Track(reporting.Copy, blob.GCS, &err)()

	src := b.bucket.Object(srcName)
	dst := b.bucket.Object(destName)

	copier := dst.CopierFrom(src)

	_, err = copier.Run(b.ctx)
	if err != nil {
		return errors.Wrapf(toStandardError(err), "could not copy from %s to %s", srcName, destName)
	}
	return nil
}

func (b *bucket) Name() string {
	return b.bucketName
}

func (b *bucket) GetDirContent(ctx context.Context, dirPath string) (dirs []string, objects []string, err error) {
	defer reporting.Track(reporting.GetDirContent, blob.GCS, &err)()

	if dirPath != "" && dirPath[len(dirPath)-1] != '/' {
		dirPath += "/"
	}
	query := &storage.Query{
		Delimiter: "/",
		Prefix:    dirPath,
		Versions:  false,
	}

	iter := b.bucket.Objects(ctx, query)
	for {
		objAttrs, iterErr := iter.Next()
		if iterErr == iterator.Done {
			break
		}
		if iterErr != nil {
			err = errors.Wrap(iterErr, "error listing objects from gcs")
			break
		}
		// If Query.Delimiter is non-empty, some of the ObjectAttrs returned by Next will
		// have a non-empty Prefix field, and a zero value for all other fields. These
		// represent prefixes.
		if objAttrs.Prefix != "" {
			dirs = append(dirs, objAttrs.Prefix)
		} else {
			objects = append(objects, objAttrs.Name)
		}
	}
	if err != nil {
		return nil, nil, errors.Wrapf(toStandardError(err), "could not list contents at %s", dirPath)
	}
	if len(objects) == 0 && len(dirs) == 0 {
		return nil, nil, errors.Newf(nil, errors.NotFound, "could not find directory %v", dirPath)
	}
	return
}

func (b *bucket) GetContentWithCommonPrefix(ctx context.Context, prefix string, maxResults int) (paths []string, err error) {
	defer reporting.Track(reporting.GetContentWithCommonPrefix, blob.GCS, &err)()
	query := &storage.Query{
		Prefix:   prefix,
		Versions: false,
	}

	var objects []string
	iter := b.bucket.Objects(ctx, query)
	for maxResults == 0 || len(objects) < maxResults {
		objAttrs, iterErr := iter.Next()
		if iterErr == iterator.Done {
			break
		}
		if iterErr != nil {
			err = errors.Wrap(iterErr, "error listing objects from gcs")
			break
		}
		objects = append(objects, objAttrs.Name)
	}
	if err != nil {
		return nil, errors.Wrapf(toStandardError(err), "Could not list contents at %s", prefix)
	}

	return objects, nil
}

func toAttributes(gcsAttr *storage.ObjectAttrs) *blob.ObjectAttributes {
	return &blob.ObjectAttributes{
		CacheControl:       gcsAttr.CacheControl,
		ContentEncoding:    gcsAttr.ContentEncoding,
		ContentDisposition: gcsAttr.ContentDisposition,
		ContentLanguage:    gcsAttr.ContentLanguage,
		ContentType:        gcsAttr.ContentType,
		Metadata:           gcsAttr.Metadata,
		MD5:                gcsAttr.MD5,
		Size:               gcsAttr.Size,
		LastModified:       gcsAttr.Updated,
	}
}

func aclToGcsRules(acl blob.ACL) string {
	// The list of ACLs is in:
	// https://cloud.google.com/storage/docs/json_api/v1/objects/insert
	switch acl {
	case blob.PUBLIC_READ:
		return "publicRead"
	case blob.AUTHENTICATED_READ:
		return "authenticatedRead"
	case blob.PRIVATE:
		fallthrough
	default:
		return "private"
	}
}

func toStandardError(err error) error {
	switch err {
	case nil:
		return nil
	case storage.ErrBucketNotExist:
		return errors.New(err, errors.NotFound, storage.ErrBucketNotExist.Error())
	case storage.ErrObjectNotExist:
		return errors.New(err, errors.NotFound, storage.ErrObjectNotExist.Error())
	default:
		return errors.New(err, errors.Internal, err.Error())
	}
}
