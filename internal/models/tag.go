package models

// Tag represents a GHL tag.
type Tag struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	LocationID string `json:"locationId"`
}

// TagsResponse is the response from GET /locations/:locationId/tags.
type TagsResponse struct {
	Tags []Tag `json:"tags"`
}

// ContactTagsRequest is the body for POST /contacts/:contactId/tags.
type ContactTagsRequest struct {
	Tags []string `json:"tags"`
}

// ContactTagsResponse is the response from tag operations.
type ContactTagsResponse struct {
	Tags []string `json:"tags"`
}
