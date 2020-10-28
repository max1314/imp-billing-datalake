// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package util

import (
	"context"
	"crypto/tls"
	"google.golang.org/api/option"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"golang.org/x/oauth2/google"

	dl "github.com/improbable/imp-billing-datalake"
	"github.com/improbable/imp-billing-datalake/errors"
	"github.com/improbable/imp-billing-datalake/internal/gcs"
	"github.com/improbable/imp-billing-datalake/internal/s3"
)

type DataLakeConfig struct {
	StorageContext context.Context
	BucketName     string
	// s3 or gcs
	StorageBackend dl.BackingStorage

	// for s3
	S3CredentialKey    string
	S3CredentialSecret string
	S3Endpoint         string //etc: s3.eu-west-2.amazonaws.com
	S3Region           string //etc: cn-northwest-1

	// for gcs
	// must be like: {"type":"service_account", "client_email":"", "private_key_id":"", "private_key":"", "token_uri":"", "project_id":""}
	GCSCredentialJsonBytes []byte // needed

	GCSEndpoint           string // custom the GCS Endpoint, for unit test
	GCSInsecureSkipVerify bool   // skip GCS tls verify, for unit test
}

func NewDataLakeFromConfig(mp *DataLakeConfig) (dl.DataLake, error) {
	if mp.StorageContext == nil {
		mp.StorageContext = context.Background()
	}

	switch mp.StorageBackend {
	case dl.S3:
		s3Creds := credentials.NewStaticCredentials(mp.S3CredentialKey, mp.S3CredentialSecret, "")
		return s3.NewDataLake(mp.StorageContext, mp.BucketName, mp.S3Endpoint, mp.S3Region, s3Creds)

	case dl.GCS:
		var opts []option.ClientOption
		if mp.GCSCredentialJsonBytes != nil {
			jwtConfig, err := google.JWTConfigFromJSON(mp.GCSCredentialJsonBytes, storage.ScopeReadWrite, storage.ScopeFullControl)
			if err != nil {
				return nil, errors.New(err, "gcp jwt config error")
			}
			opts = append(opts, option.WithTokenSource(jwtConfig.TokenSource(mp.StorageContext)))
		}
		// for unit test
		if mp.GCSEndpoint != "" {
			opts = append(opts, option.WithEndpoint(mp.GCSEndpoint))
		}
		// for unit test
		if mp.GCSInsecureSkipVerify {
			transCfg := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
			}
			httpClient := &http.Client{Transport: transCfg}
			opts = append(opts, option.WithHTTPClient(httpClient))
		}
		return gcs.NewDataLake(mp.StorageContext, mp.BucketName, opts...)

	default:
		return nil, errors.New(nil, "invalid data lake storage backend")
	}
}
