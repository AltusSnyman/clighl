package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/altusmusic/clighl/internal/models"
)

// SearchConversations searches conversations with optional filters.
func (c *Client) SearchConversations(ctx context.Context, contactID string, query string, page, limit int) (*models.ConversationsSearchResponse, error) {
	if limit <= 0 {
		limit = 20
	}

	params := url.Values{}
	params.Set("locationId", c.LocationID)
	params.Set("limit", strconv.Itoa(limit))
	if contactID != "" {
		params.Set("contactId", contactID)
	}
	if query != "" {
		params.Set("q", query)
	}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}

	var resp models.ConversationsSearchResponse
	err := c.Do(ctx, "GET", "/conversations/search?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("search conversations: %w", err)
	}
	return &resp, nil
}

// GetMessages retrieves messages for a conversation.
func (c *Client) GetMessages(ctx context.Context, conversationID string, limit int, lastMsgID string) (*models.MessagesResponse, error) {
	if limit <= 0 {
		limit = 20
	}

	params := url.Values{}
	params.Set("limit", strconv.Itoa(limit))
	if lastMsgID != "" {
		params.Set("lastMessageId", lastMsgID)
	}

	// GHL returns messages in a flexible structure — parse raw first
	var raw map[string]interface{}
	err := c.Do(ctx, "GET", "/conversations/"+conversationID+"/messages?"+params.Encode(), nil, &raw)
	if err != nil {
		return nil, fmt.Errorf("get messages: %w", err)
	}

	resp := &models.MessagesResponse{}

	// Extract messages from response — could be under "messages" key or nested
	msgData, _ := json.Marshal(raw["messages"])
	if msgData != nil {
		// Try as array first
		var msgs []models.Message
		if json.Unmarshal(msgData, &msgs) == nil {
			resp.Messages = msgs
		} else {
			// Try as object with nested array
			var wrapper map[string]interface{}
			if json.Unmarshal(msgData, &wrapper) == nil {
				if msgsArr, ok := wrapper["messages"]; ok {
					inner, _ := json.Marshal(msgsArr)
					json.Unmarshal(inner, &resp.Messages)
				}
			}
		}
	}

	if np, ok := raw["nextPage"].(bool); ok {
		resp.NextPage = np
	}
	if lm, ok := raw["lastMessageId"].(string); ok {
		resp.LastMsgId = lm
	}

	return resp, nil
}

// SendMessage sends a message in a conversation.
func (c *Client) SendMessage(ctx context.Context, req *models.SendMessageRequest) (*models.SendMessageResponse, error) {
	var resp models.SendMessageResponse
	err := c.Do(ctx, "POST", "/conversations/messages", req, &resp)
	if err != nil {
		return nil, fmt.Errorf("send message: %w", err)
	}
	return &resp, nil
}
