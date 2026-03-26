package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/altusmusic/clighl/internal/models"
)

// SearchContacts searches contacts by query string.
func (c *Client) SearchContacts(ctx context.Context, query string, page, limit int) (*models.ContactSearchResponse, error) {
	if limit <= 0 {
		limit = 20
	}
	if page <= 0 {
		page = 1
	}

	body := models.ContactSearchRequest{
		LocationID: c.LocationID,
		Query:      query,
		Page:       page,
		PageLimit:  limit,
	}

	var resp models.ContactSearchResponse
	err := c.Do(ctx, "POST", "/contacts/search", body, &resp)
	if err != nil {
		return nil, fmt.Errorf("search contacts: %w", err)
	}
	return &resp, nil
}

// GetContact retrieves a single contact by ID.
func (c *Client) GetContact(ctx context.Context, id string) (*models.Contact, error) {
	var resp models.ContactResponse
	err := c.Do(ctx, "GET", "/contacts/"+id, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get contact: %w", err)
	}
	return &resp.Contact, nil
}

// ListContacts lists contacts for the location.
func (c *Client) ListContacts(ctx context.Context, page, limit int) (*models.ContactSearchResponse, error) {
	if limit <= 0 {
		limit = 20
	}

	params := url.Values{}
	params.Set("locationId", c.LocationID)
	params.Set("limit", strconv.Itoa(limit))
	if page > 1 {
		params.Set("skip", strconv.Itoa((page-1)*limit))
	}

	var resp models.ContactSearchResponse
	err := c.Do(ctx, "GET", "/contacts/?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("list contacts: %w", err)
	}
	return &resp, nil
}

// CreateContact creates a new contact.
func (c *Client) CreateContact(ctx context.Context, req *models.ContactCreateRequest) (*models.Contact, error) {
	req.LocationID = c.LocationID
	var resp models.ContactResponse
	err := c.Do(ctx, "POST", "/contacts/", req, &resp)
	if err != nil {
		return nil, fmt.Errorf("create contact: %w", err)
	}
	return &resp.Contact, nil
}

// UpdateContact updates an existing contact.
func (c *Client) UpdateContact(ctx context.Context, id string, req *models.ContactUpdateRequest) (*models.Contact, error) {
	var resp models.ContactResponse
	err := c.Do(ctx, "PUT", "/contacts/"+id, req, &resp)
	if err != nil {
		return nil, fmt.Errorf("update contact: %w", err)
	}
	return &resp.Contact, nil
}

// UpsertContact creates or updates a contact (matched by email or phone).
func (c *Client) UpsertContact(ctx context.Context, req *models.ContactUpsertRequest) (*models.Contact, error) {
	req.LocationID = c.LocationID
	var resp models.ContactResponse
	err := c.Do(ctx, "POST", "/contacts/upsert", req, &resp)
	if err != nil {
		return nil, fmt.Errorf("upsert contact: %w", err)
	}
	return &resp.Contact, nil
}

// ListTasks returns all tasks for a contact.
func (c *Client) ListTasks(ctx context.Context, contactID string) ([]models.Task, error) {
	var resp models.TasksResponse
	err := c.Do(ctx, "GET", "/contacts/"+contactID+"/tasks", nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}
	return resp.Tasks, nil
}
