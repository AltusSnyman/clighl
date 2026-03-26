package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var calendarsSlotsCmd = &cobra.Command{
	Use:   "slots <calendar-name-or-id>",
	Short: "Show available time slots for a calendar",
	Long: `Show available booking slots for a calendar.

By default shows slots for the next 7 days.

Examples:
  clighl calendars slots "Sales Call"
  clighl calendars slots "Sales Call" --start 2026-04-01 --end 2026-04-07
  clighl calendars slots "Sales Call" --days 14 --timezone "America/New_York"`,
	Args: cobra.ExactArgs(1),
	RunE: runCalendarsSlots,
}

var (
	slotsStart    string
	slotsEnd      string
	slotsDays     int
	slotsTimezone string
)

func init() {
	calendarsCmd.AddCommand(calendarsSlotsCmd)
	calendarsSlotsCmd.Flags().StringVar(&slotsStart, "start", "", "Start date (YYYY-MM-DD, default: today)")
	calendarsSlotsCmd.Flags().StringVar(&slotsEnd, "end", "", "End date (YYYY-MM-DD, default: start + days)")
	calendarsSlotsCmd.Flags().IntVar(&slotsDays, "days", 7, "Number of days to show (default 7)")
	calendarsSlotsCmd.Flags().StringVar(&slotsTimezone, "timezone", "", "Timezone (e.g. America/New_York)")
}

func runCalendarsSlots(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)

	// Resolve calendar by name or ID
	cal, err := res.ResolveCalendar(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	// Calculate date range
	now := time.Now()
	startDate := now
	if slotsStart != "" {
		startDate, err = time.Parse("2006-01-02", slotsStart)
		if err != nil {
			return fmt.Errorf("invalid start date '%s': use YYYY-MM-DD format", slotsStart)
		}
	}

	endDate := startDate.AddDate(0, 0, slotsDays)
	if slotsEnd != "" {
		endDate, err = time.Parse("2006-01-02", slotsEnd)
		if err != nil {
			return fmt.Errorf("invalid end date '%s': use YYYY-MM-DD format", slotsEnd)
		}
	}

	// Format as UNIX ms timestamps (what GHL expects)
	startMs := fmt.Sprintf("%d", startDate.UnixMilli())
	endMs := fmt.Sprintf("%d", endDate.UnixMilli())

	fmt.Printf("Available slots for '%s' (%s to %s):\n\n",
		cal.Name, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	slots, err := client.GetFreeSlots(cmd.Context(), cal.ID, startMs, endMs, slotsTimezone)
	if err != nil {
		return err
	}

	if len(slots) == 0 {
		fmt.Println("No available slots in this date range.")
		return nil
	}

	fmt.Print(getFormatter().FormatFreeSlots(slots))
	return nil
}
