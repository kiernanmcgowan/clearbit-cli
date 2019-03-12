package cmd

import (
	"github.com/clearbit/clearbit-go/clearbit"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(autocomplete)
}

var autocomplete = &cobra.Command{
	Use:   "autocomplete <query>",
	Short: "Autocompletes a comany name",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results, resp, err := client.Autocomplete.Suggest(clearbit.AutocompleteSuggestParams{
			Query: args[0],
		})
		handleResponse(results, resp, err)
	},
}
