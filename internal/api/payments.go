package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/altusmusic/clighl/internal/models"
)

// GetOrder retrieves a payment order by ID.
func (c *Client) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)

	var resp models.OrderResponse
	err := c.Do(ctx, "GET", "/payments/orders/"+orderID+"?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get order: %w", err)
	}
	return &resp.Order, nil
}

// ListTransactions returns paginated transactions.
func (c *Client) ListTransactions(ctx context.Context, contactID string, page, limit int) (*models.TransactionsResponse, error) {
	if limit <= 0 {
		limit = 20
	}

	params := url.Values{}
	params.Set("locationId", c.LocationID)
	params.Set("limit", fmt.Sprintf("%d", limit))
	if page > 0 {
		params.Set("page", fmt.Sprintf("%d", page))
	}
	if contactID != "" {
		params.Set("contactId", contactID)
	}

	var resp models.TransactionsResponse
	err := c.Do(ctx, "GET", "/payments/transactions?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("list transactions: %w", err)
	}
	return &resp, nil
}
