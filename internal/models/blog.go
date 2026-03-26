package models

// Blog represents a GHL blog site.
type Blog struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LocationID  string `json:"locationId"`
	URL         string `json:"url"`
	DateAdded   string `json:"dateAdded"`
}

// BlogsResponse is the response from GET /blogs.
type BlogsResponse struct {
	Blogs []Blog `json:"blogs"`
}

// BlogPost represents a GHL blog post.
type BlogPost struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Slug        string   `json:"slug"`
	Status      string   `json:"status"`
	Author      string   `json:"author"`
	CategoryID  string   `json:"categoryId"`
	BlogID      string   `json:"blogId"`
	RawHTML     string   `json:"rawHTML,omitempty"`
	DateAdded   string   `json:"dateAdded"`
	DateUpdated string   `json:"dateUpdated"`
	Tags        []string `json:"tags,omitempty"`
}

// BlogPostsResponse is the response from GET /blogs/{id}/posts.
type BlogPostsResponse struct {
	Posts []BlogPost  `json:"posts"`
	Total interface{} `json:"total"`
}

// BlogPostResponse wraps a single blog post.
type BlogPostResponse struct {
	Post BlogPost `json:"data"`
}

// BlogPostCreateRequest is the body for POST /blogs/posts.
type BlogPostCreateRequest struct {
	Title      string   `json:"title"`
	BlogID     string   `json:"blogId"`
	RawHTML    string   `json:"rawHTML,omitempty"`
	Status     string   `json:"status,omitempty"`
	Author     string   `json:"author,omitempty"`
	CategoryID string   `json:"categoryId,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	Slug       string   `json:"slug,omitempty"`
	LocationID string   `json:"locationId"`
}

// BlogPostUpdateRequest is the body for PUT /blogs/posts/{id}.
type BlogPostUpdateRequest struct {
	Title      string   `json:"title,omitempty"`
	RawHTML    string   `json:"rawHTML,omitempty"`
	Status     string   `json:"status,omitempty"`
	Author     string   `json:"author,omitempty"`
	CategoryID string   `json:"categoryId,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	Slug       string   `json:"slug,omitempty"`
	BlogID     string   `json:"blogId"`
	LocationID string   `json:"locationId"`
}

// BlogAuthor represents a blog author.
type BlogAuthor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// BlogAuthorsResponse is the response from GET /blogs/authors.
type BlogAuthorsResponse struct {
	Authors []BlogAuthor `json:"authors"`
}

// BlogCategory represents a blog category.
type BlogCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// BlogCategoriesResponse is the response from GET /blogs/categories.
type BlogCategoriesResponse struct {
	Categories []BlogCategory `json:"categories"`
}

// SlugCheckResponse is the response from checking a blog URL slug.
type SlugCheckResponse struct {
	Exists    bool   `json:"exists"`
	Slug      string `json:"slug"`
	Available bool   `json:"available"`
}
