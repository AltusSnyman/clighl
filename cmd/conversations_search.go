package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var conversationsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search conversations",
	Long: `Search and filter conversations.

Examples:
  clighl conversations search
  clighl conversations search --contact "Dan"
  clighl conversations search --query "pricing"`,
	RunE: runConversationsSearch,
}

var (
	convSearchContact string
	convSearchQuery   string
	convSearchLimit   int
	convSearchPage    int
)

func init() {
	conversationsCmd.AddCommand(conversationsSearchCmd)
	conversationsSearchCmd.Flags().StringVar(&convSearchContact, "contact", "", "Filter by contact name or ID")
	conversationsSearchCmd.Flags().StringVar(&convSearchQuery, "query", "", "Search query")
	conversationsSearchCmd.Flags().IntVar(&convSearchLimit, "limit", 20, "Results per page")
	conversationsSearchCmd.Flags().IntVar(&convSearchPage, "page", 1, "Page number")
}

func runConversationsSearch(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	contactID := ""
	if convSearchContact != "" {
		res := resolver.NewResolver(client)
		contact, err := res.ResolveContact(cmd.Context(), convSearchContact)
		if err != nil {
			return err
		}
		contactID = contact.ID
	}

	resp, err := client.SearchConversations(cmd.Context(), contactID, convSearchQuery, convSearchPage, convSearchLimit)
	if err != nil {
		return err
	}

	if len(resp.Conversations) == 0 {
		fmt.Println("No conversations found.")
		return nil
	}

	fmt.Printf("Found %d conversation(s):\n\n", resp.Total)
	fmt.Print(getFormatter().FormatConversations(resp.Conversations))
	return nil
}
