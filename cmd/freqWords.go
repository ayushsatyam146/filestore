package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(freqWordsCmd)
	freqWordsCmd.Flags().StringP("order", "s", "dsc", "sort order")
	freqWordsCmd.Flags().StringP("limit", "l", "10", "limit")
}

func freqWordsCLIHandler(baseURL string) {

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

var freqWordsCmd = &cobra.Command{
	Use:   "freq-words",
	Short: "frequent words command",
	Long:  "frequent words command",
	Run: func(cmd *cobra.Command, args []string) {
		sortorder, _ := cmd.Flags().GetString("order")
		limit, _ := cmd.Flags().GetString("limit")
		baseURL := os.Getenv("BASE_URL") + "/frequentWords"
		if sortorder != "" {
			baseURL = baseURL + "?sortOrder=" + sortorder
		}
		if limit != "" {
			baseURL = baseURL + "&limit=" + limit
		}
		freqWordsCLIHandler(baseURL)
	},
}
