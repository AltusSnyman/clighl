package models

import "fmt"

// Conversation represents a GHL conversation.
// GHL returns mixed types for dates (string or number), so we use interface{}.
type Conversation struct {
	ID            string      `json:"id"`
	ContactID     string      `json:"contactId"`
	LocationID    string      `json:"locationId"`
	Type          string      `json:"type"`
	LastMessageAt interface{} `json:"lastMessageDate"`
	FullName      string      `json:"fullName"`
	ContactName   string      `json:"contactName"`
	Email         string      `json:"email"`
	Phone         string      `json:"phone"`
	Unread        interface{} `json:"unreadCount,omitempty"`
	DateAdded     interface{} `json:"dateAdded"`
	DateUpdated   interface{} `json:"dateUpdated"`
}

// LastMessageDate returns a string representation of the last message date.
func (c *Conversation) LastMessageDate() string {
	if c.LastMessageAt == nil {
		return ""
	}
	return fmt.Sprintf("%v", c.LastMessageAt)
}

// ConversationsSearchResponse is the response from GET /conversations/search.
type ConversationsSearchResponse struct {
	Conversations []Conversation `json:"conversations"`
	Total         int            `json:"total"`
}

// Message represents a message in a conversation.
type Message struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	Direction      string `json:"direction"`
	Status         string `json:"status"`
	Body           string `json:"body"`
	ContactID      string `json:"contactId"`
	ConversationID string `json:"conversationId"`
	DateAdded      string `json:"dateAdded"`
	ContentType    string `json:"contentType"`
}

// MessagesResponse is the response from GET /conversations/{id}/messages.
type MessagesResponse struct {
	Messages   []Message `json:"messages"`
	NextPage   bool      `json:"nextPage,omitempty"`
	LastMsgId  string    `json:"lastMessageId,omitempty"`
}

// SendMessageRequest is the body for POST /conversations/messages.
type SendMessageRequest struct {
	Type           string `json:"type"`
	ContactID      string `json:"contactId"`
	Message        string `json:"message,omitempty"`
	Subject        string `json:"subject,omitempty"`
	HTML           string `json:"html,omitempty"`
	ConversationID string `json:"conversationId,omitempty"`
	EmailFrom      string `json:"emailFrom,omitempty"`
}

// SendMessageResponse wraps the API response from sending a message.
type SendMessageResponse struct {
	ConversationID string  `json:"conversationId"`
	MessageID      string  `json:"messageId"`
	Message        Message `json:"message"`
}
