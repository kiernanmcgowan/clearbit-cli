package cmd

import (
	"github.com/clearbit/clearbit-go/clearbit"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(discovery)
}

var discovery = &cobra.Command{
	Use:   "discovery <query>",
	Short: "Discovery searches for companies w/ complex queries",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results, resp, err := client.Discovery.Search(clearbit.DiscoverySearchParams{
			Query: args[0],
		})
		handleResponse(results, resp, err)
	},
}
