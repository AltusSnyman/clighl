package api

import (
	"context"
	"fmt"

	"github.com/altusmusic/clighl/internal/models"
)

// ListNotes returns all notes for a contact.
func (c *Client) ListNotes(ctx context.Context, contactID string) ([]models.Note, error) {
	var resp models.NotesResponse
	err := c.Do(ctx, "GET", "/contacts/"+contactID+"/notes", nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("list notes: %w", err)
	}
	return resp.Notes, nil
}

// CreateNote adds a note to a contact.
func (c *Client) CreateNote(ctx context.Context, contactID string, body string) (*models.Note, error) {
	req := &models.NoteCreateRequest{Body: body}
	var resp models.NoteResponse
	err := c.Do(ctx, "POST", "/contacts/"+contactID+"/notes", req, &resp)
	if err != nil {
		return nil, fmt.Errorf("create note: %w", err)
	}
	return &resp.Note, nil
}

// UpdateNote updates an existing note.
func (c *Client) UpdateNote(ctx context.Context, contactID, noteID string, body string) (*models.Note, error) {
	req := &models.NoteCreateRequest{Body: body}
	var resp models.NoteResponse
	err := c.Do(ctx, "PUT", "/contacts/"+contactID+"/notes/"+noteID, req, &resp)
	if err != nil {
		return nil, fmt.Errorf("update note: %w", err)
	}
	return &resp.Note, nil
}

// DeleteNote removes a note from a contact.
func (c *Client) DeleteNote(ctx context.Context, contactID, noteID string) error {
	err := c.Do(ctx, "DELETE", "/contacts/"+contactID+"/notes/"+noteID, nil, nil)
	if err != nil {
		return fmt.Errorf("delete note: %w", err)
	}
	return nil
}
