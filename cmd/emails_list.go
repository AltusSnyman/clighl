package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var emailsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List email templates",
	RunE:  runEmailsList,
}

func init() {
	emailsCmd.AddCommand(emailsListCmd)
}

func runEmailsList(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	templates, err := client.GetEmailTemplates(cmd.Context())
	if err != nil {
		return err
	}

	if len(templates) == 0 {
		fmt.Println("No email templates found.")
		return nil
	}

	fmt.Print(getFormatter().FormatEmailTemplates(templates))
	return nil
}
