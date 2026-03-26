package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/altusmusic/clighl/internal/models"
)

// ListCalendars returns all calendars for the location.
func (c *Client) ListCalendars(ctx context.Context) ([]models.Calendar, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)

	var resp models.CalendarsResponse
	err := c.Do(ctx, "GET", "/calendars/?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("list calendars: %w", err)
	}
	return resp.Calendars, nil
}

// GetCalendar returns a single calendar by ID.
func (c *Client) GetCalendar(ctx context.Context, calendarID string) (*models.Calendar, error) {
	var resp models.CalendarResponse
	err := c.Do(ctx, "GET", "/calendars/"+calendarID, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get calendar: %w", err)
	}
	return &resp.Calendar, nil
}

// GetFreeSlots returns available time slots for a calendar.
// startDate and endDate should be in UNIX timestamp (milliseconds) or ISO 8601 format.
func (c *Client) GetFreeSlots(ctx context.Context, calendarID, startDate, endDate, timezone string) (map[string][]string, error) {
	params := url.Values{}
	params.Set("startDate", startDate)
	params.Set("endDate", endDate)
	if timezone != "" {
		params.Set("timezone", timezone)
	}

	// The API returns a flexible structure — try multiple shapes
	var raw map[string]interface{}
	err := c.Do(ctx, "GET", "/calendars/"+calendarID+"/free-slots?"+params.Encode(), nil, &raw)
	if err != nil {
		return nil, fmt.Errorf("get free slots: %w", err)
	}

	result := make(map[string][]string)
	for date, val := range raw {
		if date == "_metadata" || date == "traceId" {
			continue
		}
		switch v := val.(type) {
		case map[string]interface{}:
			// Nested structure like {"slots": [...]}
			if slots, ok := v["slots"]; ok {
				result[date] = toStringSlice(slots)
			}
		case []interface{}:
			result[date] = toStringSlice(v)
		}
	}

	return result, nil
}

// CreateAppointment books a new appointment.
func (c *Client) CreateAppointment(ctx context.Context, req *models.AppointmentCreateRequest) (*models.Appointment, error) {
	req.LocationID = c.LocationID

	// Try multiple response shapes — GHL API is inconsistent
	var raw map[string]interface{}
	err := c.Do(ctx, "POST", "/calendars/events/appointments", req, &raw)
	if err != nil {
		return nil, fmt.Errorf("create appointment: %w", err)
	}

	// Build appointment from whatever fields we get
	appt := &models.Appointment{
		CalendarID:       req.CalendarID,
		ContactID:        req.ContactID,
		SelectedSlot:     req.SelectedSlot,
		SelectedTimezone: req.SelectedTimezone,
		Notes:            req.Notes,
	}

	// Extract ID from response (could be nested under different keys)
	if id, ok := raw["id"].(string); ok {
		appt.ID = id
	}
	if event, ok := raw["event"].(map[string]interface{}); ok {
		if id, ok := event["id"].(string); ok {
			appt.ID = id
		}
		if status, ok := event["appointmentStatus"].(string); ok {
			appt.AppointmentStatus = status
		}
	}
	if status, ok := raw["appointmentStatus"].(string); ok {
		appt.AppointmentStatus = status
	}

	return appt, nil
}

// GetAppointment returns an appointment by ID.
func (c *Client) GetAppointment(ctx context.Context, eventID string) (*models.Appointment, error) {
	var resp models.AppointmentResponse
	err := c.Do(ctx, "GET", "/calendars/events/appointments/"+eventID, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get appointment: %w", err)
	}
	return &resp.Appointment, nil
}

// UpdateAppointment updates an existing appointment.
func (c *Client) UpdateAppointment(ctx context.Context, eventID string, req *models.AppointmentUpdateRequest) (*models.Appointment, error) {
	var resp models.AppointmentResponse
	err := c.Do(ctx, "PUT", "/calendars/events/appointments/"+eventID, req, &resp)
	if err != nil {
		return nil, fmt.Errorf("update appointment: %w", err)
	}
	return &resp.Appointment, nil
}

// CancelAppointment cancels/deletes an appointment.
func (c *Client) CancelAppointment(ctx context.Context, eventID string) error {
	err := c.Do(ctx, "DELETE", "/calendars/events/"+eventID, nil, nil)
	if err != nil {
		return fmt.Errorf("cancel appointment: %w", err)
	}
	return nil
}

// GetCalendarEvents retrieves events for a calendar, user, or group.
func (c *Client) GetCalendarEvents(ctx context.Context, calendarID string, startTime, endTime string) ([]models.Appointment, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)
	if calendarID != "" {
		params.Set("calendarId", calendarID)
	}
	if startTime != "" {
		params.Set("startTime", startTime)
	}
	if endTime != "" {
		params.Set("endTime", endTime)
	}

	var resp struct {
		Events []models.Appointment `json:"events"`
	}
	err := c.Do(ctx, "GET", "/calendars/events?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get calendar events: %w", err)
	}
	return resp.Events, nil
}

// GetAppointmentNotes retrieves notes for a specific appointment.
func (c *Client) GetAppointmentNotes(ctx context.Context, appointmentID string) ([]models.Note, error) {
	var resp models.NotesResponse
	err := c.Do(ctx, "GET", "/calendars/events/appointments/"+appointmentID+"/notes", nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get appointment notes: %w", err)
	}
	return resp.Notes, nil
}

func toStringSlice(v interface{}) []string {
	arr, ok := v.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, 0, len(arr))
	for _, item := range arr {
		switch s := item.(type) {
		case string:
			result = append(result, s)
		case map[string]interface{}:
			if slot, ok := s["slot"].(string); ok {
				result = append(result, slot)
			}
		}
	}
	return result
}
