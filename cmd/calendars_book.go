package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
	"github.com/altusmusic/clighl/internal/resolver"
)

var calendarsBookCmd = &cobra.Command{
	Use:   "book",
	Short: "Book an appointment on a calendar",
	Long: `Book a contact into a calendar at a specific time slot.

The slot should be an ISO 8601 datetime string. Use 'clighl calendars slots' to find available times.

Examples:
  clighl calendars book --calendar "Sales Call" --contact "Dan" --slot "2026-04-01T10:00:00-04:00" --timezone "America/New_York"
  clighl calendars book --calendar "Onboarding" --contact "john@example.com" --slot "2026-04-02T14:30:00Z" --timezone "UTC"`,
	RunE: runCalendarsBook,
}

var (
	bookCalendar string
	bookContact  string
	bookSlot     string
	bookTimezone string
	bookTitle    string
	bookNotes    string
)

func init() {
	calendarsCmd.AddCommand(calendarsBookCmd)
	calendarsBookCmd.Flags().StringVar(&bookCalendar, "calendar", "", "Calendar name or ID (required)")
	calendarsBookCmd.Flags().StringVar(&bookContact, "contact", "", "Contact name or email (required)")
	calendarsBookCmd.Flags().StringVar(&bookSlot, "slot", "", "Time slot in ISO 8601 format (required)")
	calendarsBookCmd.Flags().StringVar(&bookTimezone, "timezone", "UTC", "Timezone (e.g. America/New_York)")
	calendarsBookCmd.Flags().StringVar(&bookTitle, "title", "", "Appointment title (optional)")
	calendarsBookCmd.Flags().StringVar(&bookNotes, "notes", "", "Appointment notes (optional)")
	calendarsBookCmd.MarkFlagRequired("calendar")
	calendarsBookCmd.MarkFlagRequired("contact")
	calendarsBookCmd.MarkFlagRequired("slot")
}

func runCalendarsBook(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)

	// Resolve calendar
	cal, err := res.ResolveCalendar(cmd.Context(), bookCalendar)
	if err != nil {
		return err
	}

	// Resolve contact
	contact, err := res.ResolveContact(cmd.Context(), bookContact)
	if err != nil {
		return err
	}

	firstName := contact.FirstName
	lastName := contact.LastName
	if firstName == "" {
		firstName = contact.DisplayName()
	}

	req := &models.AppointmentCreateRequest{
		CalendarID:       cal.ID,
		ContactID:        contact.ID,
		FirstName:        firstName,
		LastName:         lastName,
		Email:            contact.Email,
		Phone:            contact.Phone,
		SelectedSlot:     bookSlot,
		SelectedTimezone: bookTimezone,
		Title:            bookTitle,
		Notes:            bookNotes,
	}

	appointment, err := client.CreateAppointment(cmd.Context(), req)
	if err != nil {
		return err
	}

	fmt.Printf("Appointment booked: %s on %s at %s\n\n",
		contact.DisplayName(), cal.Name, bookSlot)
	fmt.Print(getFormatter().FormatAppointment(appointment))
	return nil
}
