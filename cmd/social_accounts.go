package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var socialAccountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "List social media accounts",
	RunE:  runSocialAccounts,
}

func init() {
	socialCmd.AddCommand(socialAccountsCmd)
}

func runSocialAccounts(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	accounts, err := client.GetSocialAccounts(cmd.Context())
	if err != nil {
		return err
	}

	if len(accounts) == 0 {
		fmt.Println("No social media accounts found.")
		return nil
	}

	fmt.Print(getFormatter().FormatSocialAccounts(accounts))
	return nil
}
