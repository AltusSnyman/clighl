package api

import (
	"context"
	"fmt"

	"github.com/altusmusic/clighl/internal/models"
)

// ListTags returns all tags for the location.
func (c *Client) ListTags(ctx context.Context) ([]models.Tag, error) {
	var resp models.TagsResponse
	err := c.Do(ctx, "GET", "/locations/"+c.LocationID+"/tags", nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("list tags: %w", err)
	}
	return resp.Tags, nil
}

// AddContactTags adds tag(s) to a contact.
func (c *Client) AddContactTags(ctx context.Context, contactID string, tags []string) ([]string, error) {
	req := &models.ContactTagsRequest{Tags: tags}
	var resp models.ContactTagsResponse
	err := c.Do(ctx, "POST", "/contacts/"+contactID+"/tags", req, &resp)
	if err != nil {
		return nil, fmt.Errorf("add tags: %w", err)
	}
	return resp.Tags, nil
}

// RemoveContactTags removes tag(s) from a contact.
func (c *Client) RemoveContactTags(ctx context.Context, contactID string, tags []string) error {
	req := &models.ContactTagsRequest{Tags: tags}
	err := c.Do(ctx, "DELETE", "/contacts/"+contactID+"/tags", req, nil)
	if err != nil {
		return fmt.Errorf("remove tags: %w", err)
	}
	return nil
}

// DeleteTag removes a tag from the location entirely.
func (c *Client) DeleteTag(ctx context.Context, tagID string) error {
	err := c.Do(ctx, "DELETE", "/locations/"+c.LocationID+"/tags/"+tagID, nil, nil)
	if err != nil {
		return fmt.Errorf("delete tag: %w", err)
	}
	return nil
}
