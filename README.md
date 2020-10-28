# imp-billing-datalake

## Summary
```
A datalake library can provider you the abillity of configuring underlying storage backend 
and perform many actions (like List, Write, Read, Copy, Delete, Exists) on it. So basicly 
it allows you to store and fetch data from a datalake by using it's interface. 
```

## Limitation
```
// the max size of uploading(DataLake.NewWriter) file is 500MB
MaxBufferSize=500MB
```

## for Test
```bash
// At begin, you need to startup a fake-gcs-sever by run
mkdir -p examples/data/sample-bucket && cat > examples/data/sample-bucket/myfile.txt <<< "Some text here."
docker run -d --name fake-gcs-server -p 4443:4443 -v ${PWD}/examples/data:/data fsouza/fake-gcs-server

// then, you can run 
go test ./tests
```
