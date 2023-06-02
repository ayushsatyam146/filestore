package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(wcCmd)
}

func wordCountCLIHandler(baseURL string) {

	url := baseURL
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

var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "word count command",
	Long:  "word count command",
	Run: func(cmd *cobra.Command, args []string) {
		baseURL := "http://localhost:8080/wordcount"
		wordCountCLIHandler(baseURL)
	},
}
