package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/altusmusic/clighl/internal/models"
)

// GetSocialAccounts returns social media accounts for the location.
func (c *Client) GetSocialAccounts(ctx context.Context) ([]models.SocialAccount, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)

	var resp models.SocialAccountsResponse
	err := c.Do(ctx, "GET", "/social-media-posting/oauth/"+c.LocationID+"/accounts?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get social accounts: %w", err)
	}
	return resp.Accounts, nil
}

// GetSocialStats returns statistics for social media accounts.
func (c *Client) GetSocialStats(ctx context.Context, accountIDs []string) (*models.SocialStatsResponse, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)
	for _, id := range accountIDs {
		params.Add("accountId", id)
	}

	var resp models.SocialStatsResponse
	err := c.Do(ctx, "GET", "/social-media-posting/"+c.LocationID+"/statistics?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get social stats: %w", err)
	}
	return &resp, nil
}

// GetSocialPosts returns social media posts.
func (c *Client) GetSocialPosts(ctx context.Context, accountID string, limit, page int) (*models.SocialPostsResponse, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)
	if accountID != "" {
		params.Set("accountId", accountID)
	}
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	if page > 0 {
		params.Set("page", fmt.Sprintf("%d", page))
	}

	var resp models.SocialPostsResponse
	err := c.Do(ctx, "GET", "/social-media-posting/"+c.LocationID+"/posts?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get social posts: %w", err)
	}
	return &resp, nil
}

// GetSocialPost returns a single social media post.
func (c *Client) GetSocialPost(ctx context.Context, postID string) (*models.SocialPost, error) {
	var resp models.SocialPostResponse
	err := c.Do(ctx, "GET", "/social-media-posting/"+c.LocationID+"/posts/"+postID, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get social post: %w", err)
	}
	return &resp.Post, nil
}

// CreateSocialPost creates a new social media post.
func (c *Client) CreateSocialPost(ctx context.Context, req *models.SocialPostCreateRequest) (*models.SocialPost, error) {
	req.LocationID = c.LocationID

	var resp models.SocialPostResponse
	err := c.Do(ctx, "POST", "/social-media-posting/"+c.LocationID+"/posts", req, &resp)
	if err != nil {
		return nil, fmt.Errorf("create social post: %w", err)
	}
	return &resp.Post, nil
}

// UpdateSocialPost updates a social media post.
func (c *Client) UpdateSocialPost(ctx context.Context, postID string, req *models.SocialPostUpdateRequest) (*models.SocialPost, error) {
	req.LocationID = c.LocationID

	var resp models.SocialPostResponse
	err := c.Do(ctx, "PUT", "/social-media-posting/"+c.LocationID+"/posts/"+postID, req, &resp)
	if err != nil {
		return nil, fmt.Errorf("update social post: %w", err)
	}
	return &resp.Post, nil
}
