package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCmd)
}

func listCLIHandler() {
	url := os.Getenv("BASE_URL") + "/list"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "listing command",
	Long:  "listing command",
	Run: func(cmd *cobra.Command, args []string) {
		listCLIHandler()
	},
}
