package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var conversationsMessagesCmd = &cobra.Command{
	Use:   "messages <conversation-id>",
	Short: "Get messages for a conversation",
	Args:  cobra.ExactArgs(1),
	RunE:  runConversationsMessages,
}

var (
	msgLimit     int
	msgLastMsgID string
)

func init() {
	conversationsCmd.AddCommand(conversationsMessagesCmd)
	conversationsMessagesCmd.Flags().IntVar(&msgLimit, "limit", 20, "Number of messages")
	conversationsMessagesCmd.Flags().StringVar(&msgLastMsgID, "after", "", "Load messages after this message ID (pagination)")
}

func runConversationsMessages(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	resp, err := client.GetMessages(cmd.Context(), args[0], msgLimit, msgLastMsgID)
	if err != nil {
		return err
	}

	if len(resp.Messages) == 0 {
		fmt.Println("No messages found.")
		return nil
	}

	fmt.Print(getFormatter().FormatMessages(resp.Messages))

	if resp.NextPage {
		lastID := resp.Messages[len(resp.Messages)-1].ID
		fmt.Printf("\nMore messages available. Use --after %s to load next page.\n", lastID)
	}
	return nil
}
