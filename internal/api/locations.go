package api

import (
	"context"
	"fmt"

	"github.com/altusmusic/clighl/internal/models"
)

// GetLocation retrieves location details.
func (c *Client) GetLocation(ctx context.Context) (*models.Location, error) {
	var resp models.LocationResponse
	err := c.Do(ctx, "GET", "/locations/"+c.LocationID, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get location: %w", err)
	}
	return &resp.Location, nil
}

// GetCustomFields retrieves custom field definitions for the location.
func (c *Client) GetCustomFields(ctx context.Context) ([]models.LocationCustomField, error) {
	var resp models.LocationCustomFieldsResponse
	err := c.Do(ctx, "GET", "/locations/"+c.LocationID+"/customFields", nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get custom fields: %w", err)
	}
	return resp.LocationCustomFields, nil
}
