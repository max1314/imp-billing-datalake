// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"imp-billing-datalake/cmd/imp-billing/internal/cmd"
	"imp-billing-datalake/lib/log"
	"imp-billing-datalake/lib/sharedflags"

	"github.com/spf13/cobra"
)

func main() {
	logger := log.New()
	defer logger.Close()

	executableName := filepath.Base(os.Args[0])
	rootCmd := &cobra.Command{
		Use:   executableName,
		Short: fmt.Sprintf("%s generates superset importable resources.", executableName),
	}
	rootCmd.AddCommand(cmd.DataLakeCmd(logger))
	sharedflags.AddToCommand(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		logger.WithError(err).Error("fail to execute")
	}
}
