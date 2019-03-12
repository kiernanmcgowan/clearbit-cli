package cmd

import (
	"github.com/clearbit/clearbit-go/clearbit"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(revealCmd)
}

var revealCmd = &cobra.Command{
	Use:   "reveal <ip-address>",
	Short: "Looks up a company by IP address",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results, resp, err := client.Reveal.Find(clearbit.RevealFindParams{
			IP: args[0],
		})
		handleResponse(results, resp, err)
	},
}
