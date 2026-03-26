package models

// Order represents a GHL payment order.
type Order struct {
	ID          string      `json:"_id"`
	Name        string      `json:"name"`
	Status      string      `json:"status"`
	Amount      interface{} `json:"amount"`
	Currency    string      `json:"currency"`
	ContactID   string      `json:"contactId"`
	ContactName string      `json:"contactName"`
	DateAdded   interface{} `json:"createdAt"`
}

// OrderResponse wraps a single order.
type OrderResponse struct {
	Order Order `json:"data"`
}

// Transaction represents a GHL payment transaction.
type Transaction struct {
	ID          string      `json:"_id"`
	Amount      interface{} `json:"amount"`
	Currency    string      `json:"currency"`
	Status      string      `json:"status"`
	Type        string      `json:"type"`
	ContactID   string      `json:"contactId"`
	ContactName string      `json:"contactName"`
	DateAdded   interface{} `json:"createdAt"`
}

// TransactionsResponse is the response from GET /payments/transactions.
type TransactionsResponse struct {
	Transactions []Transaction `json:"data"`
	Total        interface{}   `json:"totalCount"`
}
