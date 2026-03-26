package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var calendarsNotesCmd = &cobra.Command{
	Use:   "notes <appointment-id>",
	Short: "Get notes for an appointment",
	Args:  cobra.ExactArgs(1),
	RunE:  runCalendarsNotes,
}

func init() {
	calendarsCmd.AddCommand(calendarsNotesCmd)
}

func runCalendarsNotes(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	notes, err := client.GetAppointmentNotes(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	if len(notes) == 0 {
		fmt.Println("No notes found for this appointment.")
		return nil
	}

	fmt.Printf("Appointment notes (%d):\n\n", len(notes))
	fmt.Print(getFormatter().FormatNotes(notes))
	return nil
}
