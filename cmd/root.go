package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/clearbit/clearbit-go/clearbit"
	"github.com/spf13/cobra"
)

var client *clearbit.Client

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetHandler(cli.Default)

	cbAPIKey := os.Getenv("CLEARBIT_KEY")
	if cbAPIKey == "" {
		log.Fatal("envar CLEARBIT_KEY is blank")
	}
	client = clearbit.NewClient(clearbit.WithAPIKey(cbAPIKey))
}

var rootCmd = &cobra.Command{
	Use: "clearbit",
	Run: func(cmd *cobra.Command, args []string) {
		// default help
		cmd.Help()
		os.Exit(0)
	},
}

func handleResponse(results interface{}, resp *http.Response, err error) {
	if err != nil {
		fmt.Println(err)
		log.WithError(err).Fatal("Failed to make clearbit call")
	} else if resp.StatusCode != http.StatusOK {
		log.Infof("StatusCode=%d", resp.StatusCode)
		b, _ := ioutil.ReadAll(resp.Body)
		log.Info(string(b))
	} else {
		printResult(results)
	}
}

func printResult(i interface{}) {
	// going to assume that results of Marshalable since they were Unmarshaled in response
	b, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(b))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
