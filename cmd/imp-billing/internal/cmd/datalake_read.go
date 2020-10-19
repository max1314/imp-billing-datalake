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

func ReadDataLakeFileCmd(logger *log.Logger) *cobra.Command {
	var srcPath, dstPath string
	cmd := &cobra.Command{
		Use:   "read",
		Short: "Read files from DataLake.",
		Run: func(cmd *cobra.Command, args []string) {
			arr := strings.Split(srcPath, "/")
			if len(arr) < 3 {
				logger.Fatal("src_path must like like  'storage-backend-name://bucket-name/path-name'")
			}

			storageClass := strings.Trim(arr[0], ":")
			bucketName := arr[2]
			filePath := strings.Join(arr[3:], "/")

			util.SetBlobBackingStorage(storageClass)
			readDataLakeFile(bucketName, filePath, dstPath)
		},
	}
	cmd.Flags().StringVar(&srcPath, "src_path", "", "src_path must like like 'storage-backend-name://bucket-name/path-name'")
	cmd.Flags().StringVar(&dstPath, "dst_path", "", "Local source file path.")
	return cmd
}

func readDataLakeFile(bucketName, srcPath, dstPath string) {
	dl, err := util.NewDataLakeFromFlag(bucketName)
	if err != nil {
		panic(err)
	}

	var rc io.ReadCloser
	rc, err = dl.NewReader(srcPath)
	if err != nil {
		panic(err)
	}

	defer rc.Close()

	var f *os.File
	f, err = os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileSize := 0
	var bytes []byte
	var n int
	for {
		p := make([]byte, 1024)
		n, err = rc.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		f.Write(p)
		fileSize += n
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
