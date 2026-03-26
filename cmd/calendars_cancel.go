package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var calendarsCancelCmd = &cobra.Command{
	Use:   "cancel <appointment-id>",
	Short: "Cancel an appointment",
	Args:  cobra.ExactArgs(1),
	RunE:  runCalendarsCancel,
}

func init() {
	calendarsCmd.AddCommand(calendarsCancelCmd)
}

func runCalendarsCancel(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	err = client.CancelAppointment(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Appointment %s cancelled.\n", args[0])
	return nil
}
