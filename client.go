// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package imp_billing_datalake

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/improbable/imp-billing-datalake/errors"
	"github.com/improbable/imp-billing-datalake/gcs"
	"github.com/improbable/imp-billing-datalake/s3"
	"golang.org/x/oauth2/google"
)

type BackingStorage string

const (
	GCS BackingStorage = "gcs"
	S3  BackingStorage = "s3"
)

type DataLake interface {
	List(dirPath string) ([]string, []string, error) // return (dirs, objects, err)
	NewReader(path string) (io.ReadCloser, error)
	NewWriter(path string) io.WriteCloser
	Copy(src string, dst string) error
	Delete(path string) error
	Exists(path string) (bool, error)
}

type DataLakeConfig struct {
	StorageContext  		context.Context
	BucketName 				string
	// s3 or gcs
	StorageBackend 			BackingStorage
	// for s3
	S3CredentialKey 		string
	S3CredentialSecret 		string
	S3Endpoint 				string //etc: s3.eu-west-2.amazonaws.com
	S3Region 				string //etc: cn-northwest-1
	// for gcs
	GcsCredentialAccount 	string
	// must be like: {"type":"service_account", "client_email":"", "private_key_id":"", "private_key":"", "token_uri":"", "project_id":""}
	GcsCredentialJsonBytes 	[]byte
}


func NewDataLakeFromConfig(mp *DataLakeConfig) (DataLake, error){
	if mp.StorageContext == nil {
		mp.StorageContext = context.Background()
	}

	switch mp.StorageBackend {
	case S3:
		s3Creds := credentials.NewStaticCredentials(mp.S3CredentialKey, mp.S3CredentialSecret, "")
		return s3.NewDataLake(mp.StorageContext, mp.BucketName, mp.S3Endpoint, mp.S3Region, s3Creds)
	case GCS:
		jwtConfig, err := google.JWTConfigFromJSON(mp.GcsCredentialJsonBytes, storage.ScopeReadWrite, storage.ScopeFullControl)
		if err != nil{
			return nil, err
		}
		return gcs.NewDataLake(mp.StorageContext, mp.BucketName, jwtConfig)
	default:
		return nil, errors.New("invalid data lake storage backend")
	}
}