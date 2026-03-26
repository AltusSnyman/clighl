package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var calendarsEventsCmd = &cobra.Command{
	Use:   "events",
	Short: "List calendar events/appointments",
	Long: `List events for a calendar within a date range.

Examples:
  clighl calendars events --calendar "t1"
  clighl calendars events --calendar "Sales Call" --start 2026-04-01 --end 2026-04-30`,
	RunE: runCalendarsEvents,
}

var (
	eventsCalendar string
	eventsStart    string
	eventsEnd      string
)

func init() {
	calendarsCmd.AddCommand(calendarsEventsCmd)
	calendarsEventsCmd.Flags().StringVar(&eventsCalendar, "calendar", "", "Calendar name or ID")
	calendarsEventsCmd.Flags().StringVar(&eventsStart, "start", "", "Start date (YYYY-MM-DD, default: today)")
	calendarsEventsCmd.Flags().StringVar(&eventsEnd, "end", "", "End date (YYYY-MM-DD, default: +30 days)")
}

func runCalendarsEvents(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	calendarID := ""
	if eventsCalendar != "" {
		res := resolver.NewResolver(client)
		cal, err := res.ResolveCalendar(cmd.Context(), eventsCalendar)
		if err != nil {
			return err
		}
		calendarID = cal.ID
	}

	now := time.Now()
	startTime := now.Format(time.RFC3339)
	endTime := now.AddDate(0, 0, 30).Format(time.RFC3339)

	if eventsStart != "" {
		t, err := time.Parse("2006-01-02", eventsStart)
		if err != nil {
			return fmt.Errorf("invalid start date: use YYYY-MM-DD format")
		}
		startTime = t.Format(time.RFC3339)
	}
	if eventsEnd != "" {
		t, err := time.Parse("2006-01-02", eventsEnd)
		if err != nil {
			return fmt.Errorf("invalid end date: use YYYY-MM-DD format")
		}
		endTime = t.Format(time.RFC3339)
	}

	events, err := client.GetCalendarEvents(cmd.Context(), calendarID, startTime, endTime)
	if err != nil {
		return err
	}

	if len(events) == 0 {
		fmt.Println("No events found in this date range.")
		return nil
	}

	fmt.Printf("Found %d event(s):\n\n", len(events))
	fmt.Print(getFormatter().FormatAppointments(events))
	return nil
}
