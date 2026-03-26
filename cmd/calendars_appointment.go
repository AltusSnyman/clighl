package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var calendarsAppointmentCmd = &cobra.Command{
	Use:   "appointment <id>",
	Short: "Get appointment details",
	Args:  cobra.ExactArgs(1),
	RunE:  runCalendarsAppointment,
}

func init() {
	calendarsCmd.AddCommand(calendarsAppointmentCmd)
}

func runCalendarsAppointment(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	appointment, err := client.GetAppointment(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	fmt.Print(getFormatter().FormatAppointment(appointment))
	return nil
}
