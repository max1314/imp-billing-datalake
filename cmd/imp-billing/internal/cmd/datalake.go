// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package cmd

import (
	"imp-billing-datalake/lib/log"

	"github.com/spf13/cobra"
)

func DataLakeCmd(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "datalake",
		Short: "Operate DataLake.",
	}
	cmd.AddCommand(ListDataLakeDirCmd(logger))
	cmd.AddCommand(ReadDataLakeFileCmd(logger))
	cmd.AddCommand(WriteDataLakeFileCmd(logger))
	return cmd
}
