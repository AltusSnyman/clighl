package cmd

import (
	"github.com/spf13/cobra"
)

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Manage contact notes",
	Long: `Add, list, update, and delete notes on contacts.

Notes are always attached to a contact. Use --contact to specify by name/email,
or --id for the contact ID directly.`,
}

func init() {
	rootCmd.AddCommand(notesCmd)
}
