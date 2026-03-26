package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var notesAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a note to a contact",
	Long: `Add a note to a contact.

Examples:
  clighl notes add --contact "Dan" --body "Called, will follow up next week"
  clighl notes add --contact "dan@test.com" --body "Interested in premium plan"`,
	RunE: runNotesAdd,
}

var (
	notesAddContact string
	notesAddBody    string
)

func init() {
	notesCmd.AddCommand(notesAddCmd)
	notesAddCmd.Flags().StringVar(&notesAddContact, "contact", "", "Contact name, email, or ID (required)")
	notesAddCmd.Flags().StringVar(&notesAddBody, "body", "", "Note text (required)")
	notesAddCmd.MarkFlagRequired("contact")
	notesAddCmd.MarkFlagRequired("body")
}

func runNotesAdd(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)
	contact, err := res.ResolveContact(cmd.Context(), notesAddContact)
	if err != nil {
		return err
	}

	body := strings.TrimSpace(notesAddBody)
	note, err := client.CreateNote(cmd.Context(), contact.ID, body)
	if err != nil {
		return err
	}

	fmt.Printf("Note added to %s.\n\n", contact.DisplayName())
	fmt.Print(getFormatter().FormatNote(note))
	return nil
}
