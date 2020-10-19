// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"imp-billing-datalake/lib/log"
	"imp-billing-datalake/util"

	"github.com/spf13/cobra"
)

func WriteDataLakeFileCmd(logger *log.Logger) *cobra.Command {
	var srcPath, dstPath string
	cmd := &cobra.Command{
		Use:   "write",
		Short: "Write files to DataLake.",
		Run: func(cmd *cobra.Command, args []string) {
			arr := strings.Split(dstPath, "/")
			if len(arr) < 3 {
				logger.Fatal("dst_path must like 'storage-backend-name://bucket-name/path-name'")
			}

			storageClass := strings.Trim(arr[0], ":")
			bucketName := arr[2]
			filePath := strings.Join(arr[3:], "/")

			util.SetBlobBackingStorage(storageClass)
			writeDataLakeFile(bucketName, srcPath, filePath)
		},
	}
	cmd.Flags().StringVar(&srcPath, "src_path", "", "Local source file path.")
	cmd.Flags().StringVar(&dstPath, "dst_path", "", "dst_path must like 'storage-backend-name://bucket-name/path-name'")
	return cmd
}

func writeDataLakeFile(bucketName, srcPath, dstPath string) {
	dl, err := util.NewDataLakeFromFlag(bucketName)
	if err != nil {
		panic(err)
	}

	//attr := blob.ObjectAttributes{
	//	ContentType:        "application/octet-stream",
	//	ContentEncoding:    "gzip",
	//	ContentDisposition: fmt.Sprintf(`attachment; filename="%s"`, dstPath),
	//	CacheControl:       "max-age=360",
	//	//ContentLanguage:    "english",
	//	Metadata: map[string]string{
	//		"foo": "bar",
	//	},
	//}
	//wc := dl.NewWriter(dstPath, &attr)
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

	var bytes []byte
	var n, fileSize int
	var err1, err2 error
	for {
		p := make([]byte, 1024)
		n, err1 = f.Read(p)
		if err1 != nil && err1 != io.EOF {
			panic(err1)
		}

		_, err2 = wc.Write(p)
		if err2 != nil {
			panic(err2)
		}
		fileSize += n
		bytes = append(bytes, p...)
		if err1 == io.EOF {
			break
		}
	}

	hasher := sha256.New()
	_, err = hasher.Write(bytes)
	if err != nil {
		panic(err)
	}

	mp := map[string]string{
		"fileSize": fmt.Sprintf("%v", fileSize),
		"sha256":   hex.EncodeToString(hasher.Sum(nil)),
		"srcPath":  srcPath,
		"dstPath":  dstPath,
	}
	var s []byte
	s, err = json.Marshal(mp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}
