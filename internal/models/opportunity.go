package models

// Opportunity represents a GHL opportunity.
type Opportunity struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	MonetaryValue  float64 `json:"monetaryValue"`
	PipelineID     string  `json:"pipelineId"`
	PipelineStageID string `json:"pipelineStageId"`
	Status         string  `json:"status"`
	ContactID      string  `json:"contactId"`
	LocationID     string  `json:"locationId"`
	AssignedTo     string  `json:"assignedTo"`
	DateAdded      string  `json:"dateAdded"`
	DateUpdated    string  `json:"dateUpdated"`
}

// OpportunityCreateRequest is the body for POST /opportunities/.
type OpportunityCreateRequest struct {
	PipelineID      string  `json:"pipelineId"`
	LocationID      string  `json:"locationId"`
	Name            string  `json:"name"`
	PipelineStageID string  `json:"pipelineStageId"`
	Status          string  `json:"status"`
	ContactID       string  `json:"contactId"`
	MonetaryValue   float64 `json:"monetaryValue,omitempty"`
}

// OpportunityUpdateRequest is the body for PUT /opportunities/{id}.
type OpportunityUpdateRequest struct {
	PipelineID      string  `json:"pipelineId,omitempty"`
	Name            string  `json:"name,omitempty"`
	PipelineStageID string  `json:"pipelineStageId,omitempty"`
	Status          string  `json:"status,omitempty"`
	MonetaryValue   float64 `json:"monetaryValue,omitempty"`
}

// OpportunityResponse wraps a single opportunity from the API.
type OpportunityResponse struct {
	Opportunity Opportunity `json:"opportunity"`
}

// OpportunitySearchResponse is the response from opportunity search.
type OpportunitySearchResponse struct {
	Opportunities []Opportunity `json:"opportunities"`
	Meta          Meta          `json:"meta"`
}
