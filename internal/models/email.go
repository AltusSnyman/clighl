package models

// EmailTemplate represents a GHL email template.
type EmailTemplate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Subject     string `json:"subject"`
	HTML        string `json:"html,omitempty"`
	DateAdded   string `json:"dateAdded"`
	DateUpdated string `json:"dateUpdated"`
}

// EmailTemplatesResponse is the response from GET /emails/templates.
type EmailTemplatesResponse struct {
	Templates []EmailTemplate `json:"templates"`
}

// EmailTemplateResponse wraps a single template.
type EmailTemplateResponse struct {
	Template EmailTemplate `json:"template"`
}

// EmailTemplateCreateRequest is the body for POST /emails/templates.
type EmailTemplateCreateRequest struct {
	Name       string `json:"name"`
	Subject    string `json:"subject,omitempty"`
	HTML       string `json:"html,omitempty"`
	LocationID string `json:"locationId"`
}
