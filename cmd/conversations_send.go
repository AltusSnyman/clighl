package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
	"github.com/altusmusic/clighl/internal/resolver"
)

var conversationsSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a contact",
	Long: `Send an SMS, email, or other message to a contact.

Examples:
  clighl conversations send --contact "Dan" --type SMS --message "Hey Dan, following up!"
  clighl conversations send --contact "dan@test.com" --type Email --message "Hello" --subject "Follow up"`,
	RunE: runConversationsSend,
}

var (
	sendContact string
	sendType    string
	sendMessage string
	sendSubject string
	sendConvID  string
)

func init() {
	conversationsCmd.AddCommand(conversationsSendCmd)
	conversationsSendCmd.Flags().StringVar(&sendContact, "contact", "", "Contact name, email, or ID (required)")
	conversationsSendCmd.Flags().StringVar(&sendType, "type", "SMS", "Message type: SMS, Email, WhatsApp, etc.")
	conversationsSendCmd.Flags().StringVar(&sendMessage, "message", "", "Message body (required)")
	conversationsSendCmd.Flags().StringVar(&sendSubject, "subject", "", "Email subject (for Email type)")
	conversationsSendCmd.Flags().StringVar(&sendConvID, "conversation", "", "Existing conversation ID (optional)")
	conversationsSendCmd.MarkFlagRequired("contact")
	conversationsSendCmd.MarkFlagRequired("message")
}

func runConversationsSend(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)
	contact, err := res.ResolveContact(cmd.Context(), sendContact)
	if err != nil {
		return err
	}

	req := &models.SendMessageRequest{
		Type:           sendType,
		ContactID:      contact.ID,
		ConversationID: sendConvID,
	}

	if sendType == "Email" || sendType == "email" {
		// HighLevel expects HTML/body content for email sends. Mirror the content into
		// both fields to be tolerant of API-side validation differences.
		req.Message = sendMessage
		req.HTML = sendMessage
		req.Subject = sendSubject
	} else {
		req.Message = sendMessage
		req.Subject = sendSubject
	}

	resp, err := client.SendMessage(cmd.Context(), req)
	if err != nil {
		return err
	}

	fmt.Printf("Message sent to %s via %s.\n", contact.DisplayName(), sendType)
	fmt.Printf("Conversation ID: %s\n", resp.ConversationID)
	if resp.MessageID != "" {
		fmt.Printf("Message ID:      %s\n", resp.MessageID)
	}
	return nil
}
