# imp-billing-datalake

## Config StorageBackend
- for S3
```golang
// use following golang code to configure s3 credentials
import datalake "github.com/improbable/imp-billing-datalake" 
mp := datalake.DataLakeConfig{
    BucketName: "your-s3-bucketname",
    StorageBackend: "s3",
    S3CredentialKey: "xxx",
    S3CredentialSecret: "xxx"
    S3Endpoint: "xxx",
    S3Region: "xxx"
}
```
- for GCS
```golang
// use following golang code to configure gcs credentials
import datalake "github.com/improbable/imp-billing-datalake" 
mp := datalake.DataLakeConfig{
    BucketName: "your-gcs-bucketname",
    StorageBackend: "gcs",
    GcsCredentialAccount: "xxx",
    GcsCredentialJsonBytes: "xxx"
}
// GcsCredentialJsonBytes must include these infos in bytes:
{
    "type":"service_account", 
    "client_email":"", 
    "private_key_id":"", 
    "private_key":"", 
    "token_uri":"", 
    "project_id":""
}
```

## Interface Functions
- List
```golang
List(dirPath string) ([]string, []string, error) 
// return (dirs, objects, err)
```
- NewReader
```
NewReader(path string) (io.ReadCloser, error)
```
- NewWriter
```
NewWriter(path string) io.WriteCloser
```
- Copy
```
Copy(src string, dst string) error
```
- Delete
```
Delete(path string) error
```
- Exists
```
Exists(path string) (bool, error)
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
    mp := &datalake.DataLakeConfig{
        ...
    }

    dl, err := datalake.NewDataLakeFromConfig(mp)
    if err != nil {
        panic(err)
    }
    
    fs, err := genFileStructure(dl)
    if err != nil {
        panic(err)
    }

    //TODO with fs
}
```
