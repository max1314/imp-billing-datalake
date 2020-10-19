// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"imp-billing-datalake/lib/log"
	"imp-billing-datalake/util"

	"github.com/spf13/cobra"
)

func ListDataLakeDirCmd(logger *log.Logger) *cobra.Command {
	var bucketPath string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List dirs in DataLake.",
		Run: func(cmd *cobra.Command, args []string) {
			arr := strings.Split(bucketPath, "/")
			if len(arr) < 3 {
				logger.Fatal("bucket_path must like like 'storage-backend-name://bucket-name/path-name'")
			}

			storageClass := strings.Trim(arr[0], ":")
			bucketName := arr[2]
			filePath := strings.Join(arr[3:], "/")

			util.SetBlobBackingStorage(storageClass)
			listDataLakeDir(bucketName, filePath)
		},
	}
	cmd.Flags().StringVar(&bucketPath, "bucket_path", "", "bucket_path must like like 'storage-backend-name://bucket-name/path-name'")
	return cmd
}

func listDataLakeDir(bucketName, filePath string) {
	dl, err := util.NewDataLakeFromFlag(bucketName)
	if err != nil {
		panic(err)
	}

	var dirs, objs []string
	dirs, objs, err = dl.List(filePath)
	if err != nil {
		panic(err)
	}

	rMap := map[string][]string{"dirs": dirs, "objs": objs}
	var s []byte
	s, err = json.Marshal(rMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}
