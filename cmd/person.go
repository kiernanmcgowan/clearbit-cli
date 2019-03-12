package cmd

import (
	"github.com/clearbit/clearbit-go/clearbit"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(person)
}

var person = &cobra.Command{
	Use:   "person <email>",
	Short: "Person looks up people by email",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		results, resp, err := client.Person.Find(clearbit.PersonFindParams{
			Email: args[0],
		})
		handleResponse(results, resp, err)
	},
}
