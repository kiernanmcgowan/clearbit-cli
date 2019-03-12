package cmd

import (
	"github.com/clearbit/clearbit-go/clearbit"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(company)
}

var company = &cobra.Command{
	Use:   "company <domain>",
	Short: "Looks up a company by domain name",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results, resp, err := client.Company.Find(clearbit.CompanyFindParams{
			Domain: args[0],
		})
		handleResponse(results, resp, err)
	},
}
