package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateFileCLIHandler(args []string) {
	jsonFileHashData := createFileHashJSON(args)
	filteredFiles := filterFilesToUpload(jsonFileHashData)
	uploadFilesToServer(filteredFiles)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updating files command",
	Long:  "updating files command",
	Run: func(cmd *cobra.Command, args []string) {
		updateFileCLIHandler(args)
	},
}


