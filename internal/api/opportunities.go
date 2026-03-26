package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/altusmusic/clighl/internal/models"
)

// SearchOpportunities searches opportunities with optional filters.
func (c *Client) SearchOpportunities(ctx context.Context, pipelineID string, contactID string, page, limit int) (*models.OpportunitySearchResponse, error) {
	if limit <= 0 {
		limit = 20
	}

	params := url.Values{}
	params.Set("location_id", c.LocationID)
	params.Set("limit", strconv.Itoa(limit))
	if pipelineID != "" {
		params.Set("pipeline_id", pipelineID)
	}
	if contactID != "" {
		params.Set("contact_id", contactID)
	}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	var resp models.OpportunitySearchResponse
	err := c.Do(ctx, "GET", "/opportunities/search?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("search opportunities: %w", err)
	}
	return &resp, nil
}

// GetOpportunity retrieves a single opportunity by ID.
func (c *Client) GetOpportunity(ctx context.Context, id string) (*models.Opportunity, error) {
	var resp models.OpportunityResponse
	err := c.Do(ctx, "GET", "/opportunities/"+id, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get opportunity: %w", err)
	}
	return &resp.Opportunity, nil
}

// CreateOpportunity creates a new opportunity.
func (c *Client) CreateOpportunity(ctx context.Context, req *models.OpportunityCreateRequest) (*models.Opportunity, error) {
	req.LocationID = c.LocationID
	if req.Status == "" {
		req.Status = "open"
	}

	var resp models.OpportunityResponse
	err := c.Do(ctx, "POST", "/opportunities/", req, &resp)
	if err != nil {
		return nil, fmt.Errorf("create opportunity: %w", err)
	}
	return &resp.Opportunity, nil
}

// UpdateOpportunity updates an existing opportunity.
func (c *Client) UpdateOpportunity(ctx context.Context, id string, req *models.OpportunityUpdateRequest) (*models.Opportunity, error) {
	var resp models.OpportunityResponse
	err := c.Do(ctx, "PUT", "/opportunities/"+id, req, &resp)
	if err != nil {
		return nil, fmt.Errorf("update opportunity: %w", err)
	}
	return &resp.Opportunity, nil
}
