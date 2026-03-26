package models

// Location represents a GHL location/sub-account.
type Location struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	PostalCode  string  `json:"postalCode"`
	Country     string  `json:"country"`
	Website     string  `json:"website"`
	Timezone    string  `json:"timezone"`
	LogoUrl     string  `json:"logoUrl"`
	DateAdded   string  `json:"dateAdded"`
}

// LocationResponse wraps a single location from the API.
type LocationResponse struct {
	Location Location `json:"location"`
}

// LocationCustomField represents a custom field definition for a location.
type LocationCustomField struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FieldKey   string `json:"fieldKey"`
	DataType   string `json:"dataType"`
	Placeholder string `json:"placeholder"`
	Position   int    `json:"position"`
}

// LocationCustomFieldsResponse is the response from GET /locations/{id}/customFields.
type LocationCustomFieldsResponse struct {
	LocationCustomFields []LocationCustomField `json:"customFields"`
}
