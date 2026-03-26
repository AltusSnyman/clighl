package cmd

import (
	"github.com/spf13/cobra"
)

var calendarsCmd = &cobra.Command{
	Use:     "calendars",
	Aliases: []string{"cal"},
	Short:   "Manage calendars and appointments",
	Long:    `List calendars, view available slots, and book appointments.`,
}

func init() {
	rootCmd.AddCommand(calendarsCmd)
}
