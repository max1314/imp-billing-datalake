# imp-billing-datalake

## Commands Read/Write/List from data lake
- Go To [`imp-billing`](./cmd/imp-billing/README.md) command doc

## Config StorageBackend
- for S3
```yaml
// deployment.yaml
spec:
  template:
    spec:
      containers:
      - args:
        - --datalake_blob_storage=s3
        - --s3_credentials_key_id=xxx
        - --s3_credentials_key_secret=xxx
        - --s3_endpoint=s3.cn-northwest-1.amazonaws.com.cn
        - --s3_region=cn-northwest-1
        mage: your-image:version
  ...
```
- for GCS
```yaml
// deployment.yaml
spec:
  template:
    spec:
      containers:
      - args:
        - --datalake_blob_storage=gcs
        - --gcs_credentials_account=xxx
        - --gcs_credentials_secret=xxx
        - --gcs_signed_url_expiration_time=1200 #seconds
        - --gcs_credentials_type=json
        mage: your-image:version
  ...
```

## Examples
- list files from data lake
```golang
import(
    datalake "github.com/improbable/imp-billing-datalake" 
)

func genFileStructure(dl datalake.DataLake)(map[string][]string, error){
	var dirs, objs []string
	dirs, objs, err = dl.List(filePath)
	if err != nil {
		return nil, err
	}

	rMap := map[string][]string{"dirs": dirs, "objs": objs}
	return rMap, err
}

func main(){
    dl, err := datalake.util.NewDataLakeFromFlag(bucketName)
    if err != nil {
        panic(err)
    }
    
    mp, err := genFileStructure(dl)
    if err != nil {
        panic(err)
    }

    //TODO with mp
}
```
- read a file from data lake
```golang
import(
    datalake "github.com/improbable/imp-billing-datalake" 
)

func readFileContent(dl datalake.Datalake, srcPath, dstPath string) ([]byte, error){
    rc, err := dl.NewReader(srcPath)
    if err != nil {
        panic(err)
    }
    defer rc.Close()

    f, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    
    var bytes []byte
    for {
        p := make([]byte, 1024)
        _, err = rc.Read(p)
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
        bytes = append(bytes, p)
    }
    return bytes, nil
}

func main(){
    dl, err := datalake.util.NewDataLakeFromFlag(bucketName)
    if err != nil {
        panic(err)
    }
    
    ct, err := readFileContent(dl, "s3://mybucket/filepath", "/tmp/myfile")
    if err != nil {
        panic(err)
    }

    //TODO with ct
}
```
- write a file to data lake
```golang
import(
    datalake "github.com/improbable/imp-billing-datalake" 
)

func uploadFile(dl datalake.Datalake, srcPath, dstPath string) {
	var err error
	wc := dl.NewWriter(dstPath)
	defer func() {
		err = wc.Close()
		if err != nil {
			panic(err)
		}
	}()

	var f *os.File
	f, err = os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var err1, err2 error
	for {
		p := make([]byte, 1024)
		_, err1 = f.Read(p)
		if err1 != nil && err1 != io.EOF {
			panic(err1)
		}

		_, err2 = wc.Write(p)
		if err2 != nil {
			panic(err2)
		}

		if err1 == io.EOF {
			break
		}
	}
}

func main(){
    dl, err := datalake.util.NewDataLakeFromFlag(bucketName)
    if err != nil {
        panic(err)
    }
    
    uploadFile(dl, "/tmp/myfile", "s3://mybucket/filepath")
    
    fmt.Println("success upload file!")
}
```
