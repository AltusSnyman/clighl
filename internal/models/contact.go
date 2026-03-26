package models

// Contact represents a GHL contact.
type Contact struct {
	ID             string            `json:"id"`
	FirstName      string            `json:"firstName"`
	LastName       string            `json:"lastName"`
	Name           string            `json:"name"`
	Email          string            `json:"email"`
	Phone          string            `json:"phone"`
	CompanyName    string            `json:"companyName"`
	Tags           []string          `json:"tags"`
	Source         string            `json:"source"`
	DateAdded      string            `json:"dateAdded"`
	DateUpdated    string            `json:"dateUpdated"`
	LocationID     string            `json:"locationId"`
	CustomFields   []CustomField     `json:"customFields,omitempty"`
}

// CustomField represents a custom field on a contact.
type CustomField struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// ContactSearchRequest is the body for POST /contacts/search.
type ContactSearchRequest struct {
	LocationID string `json:"locationId"`
	Query      string `json:"query,omitempty"`
	Page       int    `json:"page,omitempty"`
	PageLimit  int    `json:"pageLimit,omitempty"`
}

// ContactSearchResponse is the response from contact search.
type ContactSearchResponse struct {
	Contacts []Contact `json:"contacts"`
	Total    int       `json:"total"`
}

// ContactCreateRequest is the body for POST /contacts/.
type ContactCreateRequest struct {
	FirstName   string   `json:"firstName,omitempty"`
	LastName    string   `json:"lastName,omitempty"`
	Email       string   `json:"email,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	Name        string   `json:"name,omitempty"`
	CompanyName string   `json:"companyName,omitempty"`
	LocationID  string   `json:"locationId"`
	Tags        []string `json:"tags,omitempty"`
	Source      string   `json:"source,omitempty"`
}

// ContactUpdateRequest is the body for PUT /contacts/{id}.
type ContactUpdateRequest struct {
	FirstName   string   `json:"firstName,omitempty"`
	LastName    string   `json:"lastName,omitempty"`
	Email       string   `json:"email,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	Name        string   `json:"name,omitempty"`
	CompanyName string   `json:"companyName,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Source      string   `json:"source,omitempty"`
}

// ContactUpsertRequest is the body for POST /contacts/upsert.
type ContactUpsertRequest struct {
	FirstName   string   `json:"firstName,omitempty"`
	LastName    string   `json:"lastName,omitempty"`
	Email       string   `json:"email,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	Name        string   `json:"name,omitempty"`
	CompanyName string   `json:"companyName,omitempty"`
	LocationID  string   `json:"locationId"`
	Tags        []string `json:"tags,omitempty"`
	Source      string   `json:"source,omitempty"`
}

// ContactResponse wraps a single contact from the API.
type ContactResponse struct {
	Contact Contact `json:"contact"`
}

// DisplayName returns the best available name for display.
func (c *Contact) DisplayName() string {
	if c.Name != "" {
		return c.Name
	}
	name := c.FirstName
	if c.LastName != "" {
		if name != "" {
			name += " "
		}
		name += c.LastName
	}
	if name == "" {
		if c.Email != "" {
			return c.Email
		}
		return c.ID
	}
	return name
}
