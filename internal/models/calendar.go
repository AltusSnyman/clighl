package models

// Calendar represents a GHL calendar.
type Calendar struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LocationID  string `json:"locationId"`
	SlotDuration int   `json:"slotDuration"`
	SlotInterval int   `json:"slotInterval"`
	IsActive    bool   `json:"isActive"`
}

// CalendarsResponse is the response from GET /calendars.
type CalendarsResponse struct {
	Calendars []Calendar `json:"calendars"`
}

// CalendarResponse wraps a single calendar.
type CalendarResponse struct {
	Calendar Calendar `json:"calendar"`
}

// FreeSlotsResponse is the response from GET /calendars/:id/free-slots.
type FreeSlotsResponse map[string][]FreeSlot

// FreeSlot represents an available time slot.
type FreeSlot struct {
	Slot string `json:"slot"`
}

// FreeSlotsWrapper wraps the actual API response which nests slots under a key.
type FreeSlotsWrapper struct {
	Slots map[string][]string `json:"slots,omitempty"`
	// Some API versions return a flat structure
	AvailableSlots map[string][]FreeSlot `json:"availableSlots,omitempty"`
}

// Appointment represents a GHL calendar appointment.
type Appointment struct {
	ID               string `json:"id"`
	CalendarID       string `json:"calendarId"`
	ContactID        string `json:"contactId"`
	Title            string `json:"title"`
	Status           string `json:"status"`
	AppointmentStatus string `json:"appointmentStatus"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	SelectedSlot     string `json:"selectedSlot"`
	SelectedTimezone string `json:"selectedTimezone"`
	Notes            string `json:"notes"`
	DateAdded        string `json:"dateAdded"`
	DateUpdated      string `json:"dateUpdated"`
}

// AppointmentCreateRequest is the body for POST /calendars/events/appointments.
type AppointmentCreateRequest struct {
	CalendarID       string `json:"calendarId"`
	LocationID       string `json:"locationId"`
	ContactID        string `json:"contactId,omitempty"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email,omitempty"`
	Phone            string `json:"phone,omitempty"`
	SelectedSlot     string `json:"selectedSlot"`
	SelectedTimezone string `json:"selectedTimezone"`
	Title            string `json:"title,omitempty"`
	Notes            string `json:"notes,omitempty"`
}

// AppointmentUpdateRequest is the body for PUT /calendars/events/appointments/:id.
type AppointmentUpdateRequest struct {
	CalendarID       string `json:"calendarId,omitempty"`
	SelectedSlot     string `json:"selectedSlot,omitempty"`
	SelectedTimezone string `json:"selectedTimezone,omitempty"`
	Title            string `json:"title,omitempty"`
	Notes            string `json:"notes,omitempty"`
	Status           string `json:"status,omitempty"`
}

// AppointmentResponse wraps a single appointment from the API.
type AppointmentResponse struct {
	Appointment Appointment `json:"event"`
}
