package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var contactsSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search contacts by name, email, or phone",
	Args:  cobra.ExactArgs(1),
	RunE:  runContactsSearch,
}

var (
	searchLimit int
	searchPage  int
)

func init() {
	contactsCmd.AddCommand(contactsSearchCmd)
	contactsSearchCmd.Flags().IntVar(&searchLimit, "limit", 20, "Number of results per page")
	contactsSearchCmd.Flags().IntVar(&searchPage, "page", 1, "Page number")
}

func runContactsSearch(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	resp, err := client.SearchContacts(cmd.Context(), args[0], searchPage, searchLimit)
	if err != nil {
		return err
	}

	if len(resp.Contacts) == 0 {
		fmt.Printf("No contacts found matching '%s'\n", args[0])
		return nil
	}

	fmt.Printf("Found %d contact(s):\n\n", resp.Total)
	fmt.Print(getFormatter().FormatContacts(resp.Contacts))
	return nil
}
