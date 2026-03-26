package models

// Meta holds pagination metadata from API responses.
// GHL API returns inconsistent types for pagination fields, so all fields
// are interface{} to handle strings, numbers, and nulls gracefully.
type Meta map[string]interface{}

// SearchFilter represents a filter condition for search endpoints.
type SearchFilter struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}
