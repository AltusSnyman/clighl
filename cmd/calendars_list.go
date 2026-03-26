package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var calendarsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all calendars",
	RunE:  runCalendarsList,
}

func init() {
	calendarsCmd.AddCommand(calendarsListCmd)
}

func runCalendarsList(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	calendars, err := client.ListCalendars(cmd.Context())
	if err != nil {
		return err
	}

	if len(calendars) == 0 {
		fmt.Println("No calendars found.")
		return nil
	}

	fmt.Printf("Found %d calendar(s):\n\n", len(calendars))
	fmt.Print(getFormatter().FormatCalendars(calendars))
	return nil
}
