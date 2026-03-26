package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var notesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List notes for a contact",
	RunE:  runNotesList,
}

var notesListContact string

func init() {
	notesCmd.AddCommand(notesListCmd)
	notesListCmd.Flags().StringVar(&notesListContact, "contact", "", "Contact name, email, or ID (required)")
	notesListCmd.MarkFlagRequired("contact")
}

func runNotesList(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)
	contact, err := res.ResolveContact(cmd.Context(), notesListContact)
	if err != nil {
		return err
	}

	notes, err := client.ListNotes(cmd.Context(), contact.ID)
	if err != nil {
		return err
	}

	if len(notes) == 0 {
		fmt.Printf("No notes found for %s.\n", contact.DisplayName())
		return nil
	}

	fmt.Printf("Notes for %s (%d):\n\n", contact.DisplayName(), len(notes))
	fmt.Print(getFormatter().FormatNotes(notes))
	return nil
}
