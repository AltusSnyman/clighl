package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/altusmusic/clighl/internal/models"
)

// GetEmailTemplates returns email templates for the location.
func (c *Client) GetEmailTemplates(ctx context.Context) ([]models.EmailTemplate, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)

	var resp models.EmailTemplatesResponse
	err := c.Do(ctx, "GET", "/emails/templates?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get email templates: %w", err)
	}
	return resp.Templates, nil
}

// CreateEmailTemplate creates a new email template.
func (c *Client) CreateEmailTemplate(ctx context.Context, req *models.EmailTemplateCreateRequest) (*models.EmailTemplate, error) {
	req.LocationID = c.LocationID

	var resp models.EmailTemplateResponse
	err := c.Do(ctx, "POST", "/emails/templates", req, &resp)
	if err != nil {
		return nil, fmt.Errorf("create email template: %w", err)
	}
	return &resp.Template, nil
}
