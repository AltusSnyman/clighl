package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var contactsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contacts",
	RunE:  runContactsList,
}

var (
	listLimit int
	listPage  int
)

func init() {
	contactsCmd.AddCommand(contactsListCmd)
	contactsListCmd.Flags().IntVar(&listLimit, "limit", 20, "Number of results per page")
	contactsListCmd.Flags().IntVar(&listPage, "page", 1, "Page number")
}

func runContactsList(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	resp, err := client.ListContacts(cmd.Context(), listPage, listLimit)
	if err != nil {
		return err
	}

	if len(resp.Contacts) == 0 {
		fmt.Println("No contacts found.")
		return nil
	}

	fmt.Printf("Contacts (page %d, %d total):\n\n", listPage, resp.Total)
	fmt.Print(getFormatter().FormatContacts(resp.Contacts))
	return nil
}
