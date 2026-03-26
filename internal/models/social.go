package models

// SocialAccount represents a social media account/group.
type SocialAccount struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Avatar   string `json:"avatar"`
}

// SocialAccountsResponse is the response from GET /social-media-posting/accounts.
type SocialAccountsResponse struct {
	Accounts []SocialAccount `json:"accounts"`
}

// SocialPost represents a social media post.
type SocialPost struct {
	ID          string      `json:"id"`
	AccountID   string      `json:"accountId"`
	Status      string      `json:"status"`
	Summary     string      `json:"summary"`
	Content     string      `json:"content"`
	Platform    string      `json:"platform"`
	Type        string      `json:"type"`
	ScheduledAt string      `json:"scheduledAt"`
	PublishedAt string      `json:"publishedAt"`
	DateAdded   interface{} `json:"dateAdded"`
}

// SocialPostsResponse is the response from GET /social-media-posting/posts.
type SocialPostsResponse struct {
	Posts []SocialPost `json:"posts"`
	Total interface{}  `json:"total"`
}

// SocialPostResponse wraps a single social media post.
type SocialPostResponse struct {
	Post SocialPost `json:"post"`
}

// SocialPostCreateRequest is the body for POST /social-media-posting/posts.
type SocialPostCreateRequest struct {
	AccountIDs  []string `json:"accountIds"`
	Content     string   `json:"content"`
	Summary     string   `json:"summary,omitempty"`
	Type        string   `json:"type,omitempty"`
	ScheduledAt string   `json:"scheduledAt,omitempty"`
	LocationID  string   `json:"locationId"`
}

// SocialPostUpdateRequest is the body for PUT /social-media-posting/posts/{id}.
type SocialPostUpdateRequest struct {
	Content     string `json:"content,omitempty"`
	Summary     string `json:"summary,omitempty"`
	ScheduledAt string `json:"scheduledAt,omitempty"`
	LocationID  string `json:"locationId"`
}

// SocialStats represents social media statistics.
type SocialStats struct {
	AccountID  string      `json:"accountId"`
	Platform   string      `json:"platform"`
	Followers  interface{} `json:"followers"`
	Posts      interface{} `json:"posts"`
	Engagement interface{} `json:"engagement"`
}

// SocialStatsResponse is the response from GET /social-media-posting/statistics.
type SocialStatsResponse struct {
	Stats []SocialStats `json:"data"`
}
