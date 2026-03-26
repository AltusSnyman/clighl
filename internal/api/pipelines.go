package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/altusmusic/clighl/internal/models"
)

// ListPipelines returns all pipelines for the location.
func (c *Client) ListPipelines(ctx context.Context) ([]models.Pipeline, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)

	var resp models.PipelinesResponse
	err := c.Do(ctx, "GET", "/opportunities/pipelines?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("list pipelines: %w", err)
	}
	return resp.Pipelines, nil
}
