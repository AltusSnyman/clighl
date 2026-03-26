package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var notesDeleteCmd = &cobra.Command{
	Use:   "delete <note-id>",
	Short: "Delete a note from a contact",
	Args:  cobra.ExactArgs(1),
	RunE:  runNotesDelete,
}

var notesDeleteContact string

func init() {
	notesCmd.AddCommand(notesDeleteCmd)
	notesDeleteCmd.Flags().StringVar(&notesDeleteContact, "contact", "", "Contact name, email, or ID (required)")
	notesDeleteCmd.MarkFlagRequired("contact")
}

func runNotesDelete(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)
	contact, err := res.ResolveContact(cmd.Context(), notesDeleteContact)
	if err != nil {
		return err
	}

	err = client.DeleteNote(cmd.Context(), contact.ID, args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Note %s deleted from %s.\n", args[0], contact.DisplayName())
	return nil
}
