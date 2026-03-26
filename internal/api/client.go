package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/time/rate"

	"github.com/altusmusic/clighl/internal/config"
)

const (
	defaultBaseURL = "https://services.leadconnectorhq.com"
	defaultTimeout = 30 * time.Second
)

// Client is the GHL API client.
type Client struct {
	BaseURL     string
	AccessToken string
	LocationID  string
	APIVersion  string
	HTTPClient  *http.Client
	limiter     *rate.Limiter
}

// NewClient creates a new API client from config.
func NewClient(cfg *config.Config) *Client {
	return &Client{
		BaseURL:     defaultBaseURL,
		AccessToken: cfg.AccessToken,
		LocationID:  cfg.LocationID,
		APIVersion:  cfg.APIVersion,
		HTTPClient: &http.Client{
			Timeout: defaultTimeout,
		},
		// 10 req/s sustained, burst of 100 (conservative vs 100/10s limit)
		limiter: rate.NewLimiter(rate.Every(100*time.Millisecond), 10),
	}
}

// NewClientFromToken creates a client directly from token and location ID.
// Useful for the auth validation step before config is saved.
func NewClientFromToken(locationID, accessToken string) *Client {
	return &Client{
		BaseURL:     defaultBaseURL,
		AccessToken: accessToken,
		LocationID:  locationID,
		APIVersion:  "2021-07-28",
		HTTPClient: &http.Client{
			Timeout: defaultTimeout,
		},
		limiter: rate.NewLimiter(rate.Every(100*time.Millisecond), 10),
	}
}

// Do executes an API request with auth headers, rate limiting, and error handling.
func (c *Client) Do(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	// Wait for rate limiter
	if err := c.limiter.Wait(ctx); err != nil {
		return fmt.Errorf("rate limiter: %w", err)
	}

	// Marshal body
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	// Build request
	url := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Version", c.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Execute
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	// Handle rate limiting with single retry
	if resp.StatusCode == 429 {
		retryAfter := 1
		if v := resp.Header.Get("Retry-After"); v != "" {
			if n, err := strconv.Atoi(v); err == nil {
				retryAfter = n
			}
		}
		fmt.Fprintf(io.Discard, "Rate limited. Retrying in %d seconds...\n", retryAfter)

		select {
		case <-time.After(time.Duration(retryAfter) * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}

		return c.Do(ctx, method, path, body, result)
	}

	// Handle errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			Raw:        string(respBody),
		}
		// Try to extract message from JSON error response
		var errResp struct {
			Message string `json:"message"`
			Msg     string `json:"msg"`
		}
		if json.Unmarshal(respBody, &errResp) == nil {
			if errResp.Message != "" {
				apiErr.Message = errResp.Message
			} else if errResp.Msg != "" {
				apiErr.Message = errResp.Msg
			}
		}
		return apiErr
	}

	// Unmarshal response
	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("unmarshal response: %w", err)
		}
	}

	return nil
}
