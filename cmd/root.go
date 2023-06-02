package cmd

import (
	"github.com/joho/godotenv"
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
	godotenv.Load()
	return rootCmd.Execute()
}
