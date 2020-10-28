package tests

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
)

const GCSEndpoint = "https://storage.gcs.127.0.0.1.nip.io:4443/storage/v1/"

func FakeS3StorageTestDecorator(bucketName string, testFunc func(string, string)) {
	backend := s3mem.New()
	faker := gofakes3.New(backend)
	ts := httptest.NewServer(faker.Server())
	defer ts.Close()

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		Endpoint:         aws.String(ts.URL),
		Region:           aws.String("test"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
	cparams := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	// Create a new bucket using the CreateBucket call.
	_, err := s3Client.CreateBucket(cparams)
	if err != nil {
		panic(err)
	}

	testFunc(bucketName, ts.URL)
}

func FakeGCSStorageTestDecorator(bucketName string, testFunc func(string, string)) {
	// check if fake server and bucket exist
	bucketUrl := fmt.Sprintf("%sb/%s/o", GCSEndpoint, bucketName)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(bucketUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// This should be check when you use fake gcs http server for testing NewReader function.
	//host := os.Getenv("STORAGE_EMULATOR_HOST")
	//if host != "storage.gcs.127.0.0.1.nip.io:4443" {
	//	panic(errors.New("env:STORAGE_EMULATOR_HOST is not right! "))
	//}

	testFunc(bucketName, GCSEndpoint)
}
