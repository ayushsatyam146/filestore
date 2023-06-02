package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "filestore is a CLI",
		Long:  "filestore is a CLI",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
