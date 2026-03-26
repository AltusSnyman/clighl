package models

// Pipeline represents a GHL pipeline.
type Pipeline struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Stages     []Stage `json:"stages"`
	LocationID string  `json:"locationId"`
}

// Stage represents a stage within a pipeline.
type Stage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PipelinesResponse is the response from GET /opportunities/pipelines.
type PipelinesResponse struct {
	Pipelines []Pipeline `json:"pipelines"`
}
