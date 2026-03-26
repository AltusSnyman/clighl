package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var contactsGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get contact details by ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runContactsGet,
}

func init() {
	contactsCmd.AddCommand(contactsGetCmd)
}

func runContactsGet(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	contact, err := client.GetContact(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	fmt.Print(getFormatter().FormatContact(contact))
	return nil
}
