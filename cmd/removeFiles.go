package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

func deleteCLIHandler(args []string) {
	fileName := args[0]
	url := os.Getenv("BASE_URL") + "/deletefile"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Failed to create HTTP request:", err)
		return
	}

	q := req.URL.Query()
	q.Add("filename", fileName)
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("API request failed with status:", resp.Status)
		return
	}

	fmt.Println("File deleted Successfully")
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "removing command",
	Long:  "removing command",
	Run: func(cmd *cobra.Command, args []string) {
		deleteCLIHandler(args)
	},
}
