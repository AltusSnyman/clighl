package models

// Note represents a GHL contact note.
type Note struct {
	ID        string `json:"id"`
	Body      string `json:"body"`
	ContactID string `json:"contactId"`
	DateAdded string `json:"dateAdded"`
	DateUpdated string `json:"dateUpdated"`
}

// NoteCreateRequest is the body for POST /contacts/:contactId/notes.
type NoteCreateRequest struct {
	Body string `json:"body"`
}

// NoteResponse wraps a single note from the API.
type NoteResponse struct {
	Note Note `json:"note"`
}

// NotesResponse is the response from GET /contacts/:contactId/notes.
type NotesResponse struct {
	Notes []Note `json:"notes"`
}
