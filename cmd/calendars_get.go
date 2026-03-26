package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var calendarsGetCmd = &cobra.Command{
	Use:   "get <id-or-name>",
	Short: "Get calendar details",
	Args:  cobra.ExactArgs(1),
	RunE:  runCalendarsGet,
}

func init() {
	calendarsCmd.AddCommand(calendarsGetCmd)
}

func runCalendarsGet(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	// Try direct ID lookup first
	cal, err := client.GetCalendar(cmd.Context(), args[0])
	if err == nil {
		fmt.Print(getFormatter().FormatCalendar(cal))
		return nil
	}

	// Fall back to name search
	calendars, err := client.ListCalendars(cmd.Context())
	if err != nil {
		return err
	}

	for i := range calendars {
		if calendars[i].Name == args[0] || calendars[i].ID == args[0] {
			fmt.Print(getFormatter().FormatCalendar(&calendars[i]))
			return nil
		}
	}

	return fmt.Errorf("calendar '%s' not found", args[0])
}
