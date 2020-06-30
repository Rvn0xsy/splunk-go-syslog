package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	splunk "github.com/brittonhayes/splunk-golang/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(SearchCmd)
}

// SearchCmd is used to search splunk events
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search Splunk for events.",
	Long: `The search command is used to perform search queries via the Splunk REST API. 
Searching splunk using the CLI requires one argument of a SPL file containing your search.

e.g. splunk-go search ~/.splunk-go/searches/my-search.spl

`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		conn := splunk.Connection{
			Username: viper.GetString("SPLUNK_USERNAME"),
			Password: viper.GetString("SPLUNK_PASSWORD"),
			BaseURL:  viper.GetString("SPLUNK_URL"),
		}

		content, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}
		text := string(content)

		response, err := conn.SearchSync(text)
		if err != nil {
			log.Fatal(au.Red("Couldn't search Splunk. Ensure your configuration is correct."))
		}

		fmt.Printf(response)
	},
}
