// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Contains abstraction over the Google storage API.

package gcs

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"time"

	blob "github.com/max1314/imp-billing-datalake"
	"github.com/max1314/imp-billing-datalake/lib/errors"
	"github.com/max1314/imp-billing-datalake/internal/reporting"
	"github.com/max1314/imp-billing-datalake/lib/sharedflags"

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
func NewDataLake(ctx context.Context, bucketName string) (blob.DataLake, error) {
	return NewExplicitBucket(ctx, bucketName, *fCredsType, *fCredsSecret, *fCredsAccount)
}

// NewExplicitBucket creates a client with explicit values of creds and project
func NewExplicitBucket(ctx context.Context, bucketName string, credsType string, credsSecretPath string, credsAccount string) (blob.DataLake, error) {
	bucket := &dataLake{bucketName: bucketName, ctx: ctx}
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
func NewBucketWithTokenSource(ctx context.Context, bucketName string, source oauth2.TokenSource) (blob.DataLake, error) {
	client, err := storage.NewClient(ctx, option.WithTokenSource(source))
	if err != nil {
		return nil, err
	}

	return &dataLake{
		bucket:     client.Bucket(bucketName),
		ctx:        ctx,
		client:     client,
		bucketName: bucketName,
	}, nil
}

type dataLake struct {
	ctx               context.Context
	client            *storage.Client
	bucket            *storage.BucketHandle
	bucketName        string
	urlSigningAccount string
	urlSigningKey     []byte
}

func (b *dataLake) NewReader(path string) (r io.ReadCloser, err error) {
	defer reporting.Track(reporting.NewReader, blob.GCS, &err)()

	object := b.bucket.Object(path)
	opts := &blob.ReaderOptions{}
	if opts != nil && opts.GCSReadCompressed {
		object = object.ReadCompressed(true)
	}
	reader, err := object.NewReader(b.ctx)
	if err != nil {
		return nil, errors.Wrapf(toStandardError(err), "Could not create reader for %s", path)
	}
	return reader, nil
}

func (b *dataLake) NewWriter(path string) io.WriteCloser {
	w := b.bucket.Object(path).NewWriter(b.ctx)
	attrs := &blob.ObjectAttributes{}

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


func (b *dataLake) Exists(path string) (exists bool, err error) {
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

func (b *dataLake) Delete(path string) (err error) {
	defer reporting.Track(reporting.Delete, blob.GCS, &err)()
	err = b.bucket.Object(path).Delete(b.ctx)
	if err != nil {
		return errors.Wrapf(toStandardError(err), "could not delete %s", path)
	}
	return nil
}

// Copy calls google api to do a copy
func (b *dataLake) Copy(srcName string, destName string) (err error) {
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

func (b *dataLake) List(dirPath string) (dirs []string, objects []string, err error) {
	defer reporting.Track(reporting.GetDirContent, blob.GCS, &err)()

	if dirPath != "" && dirPath[len(dirPath)-1] != '/' {
		dirPath += "/"
	}
	query := &storage.Query{
		Delimiter: "/",
		Prefix:    dirPath,
		Versions:  false,
	}

	iter := b.bucket.Objects(b.ctx, query)
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
