package util

import (
	"cloud.google.com/go/storage"
	"context"
	"golang.org/x/oauth2/google"

	"github.com/aws/aws-sdk-go/aws/credentials"
	blob "github.com/improbable/imp-billing-datalake"
	"github.com/improbable/imp-billing-datalake/gcs"
	"github.com/improbable/imp-billing-datalake/internal/errors"
	"github.com/improbable/imp-billing-datalake/s3"
)

type DataLakeConfigMap struct {
	StorageContext  		context.Context
	BucketName 				string
	// s3 or gcs
	StorageBackend 			blob.BackingStorage
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


func NewDataLakeFromConfigMap(mp *DataLakeConfigMap) (blob.DataLake, error){
	if mp.StorageContext == nil {
		mp.StorageContext = context.Background()
	}

	switch mp.StorageBackend {
	case blob.S3:
		s3Creds := credentials.NewStaticCredentials(mp.S3CredentialKey, mp.S3CredentialSecret, "")
		return s3.NewDataLake(mp.StorageContext, mp.BucketName, mp.S3Endpoint, mp.S3Region, s3Creds)
	case blob.GCS:
		jwtConfig, err := google.JWTConfigFromJSON(mp.GcsCredentialJsonBytes, storage.ScopeReadWrite, storage.ScopeFullControl)
		if err != nil{
			return nil, err
		}
		return gcs.NewDataLake(mp.StorageContext, mp.BucketName, jwtConfig)
	default:
		return nil, errors.New("invalid data lake storage backend")
	}
}
