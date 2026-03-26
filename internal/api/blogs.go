package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/altusmusic/clighl/internal/models"
)

// ListBlogs returns all blogs for the location.
func (c *Client) ListBlogs(ctx context.Context) ([]models.Blog, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)

	var resp models.BlogsResponse
	err := c.Do(ctx, "GET", "/blogs/?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("list blogs: %w", err)
	}
	return resp.Blogs, nil
}

// GetBlogPosts returns posts for a blog.
func (c *Client) GetBlogPosts(ctx context.Context, blogID string, limit, offset int) (*models.BlogPostsResponse, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	if offset > 0 {
		params.Set("offset", fmt.Sprintf("%d", offset))
	}

	var resp models.BlogPostsResponse
	err := c.Do(ctx, "GET", "/blogs/"+blogID+"/posts?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get blog posts: %w", err)
	}
	return &resp, nil
}

// CreateBlogPost creates a new blog post.
func (c *Client) CreateBlogPost(ctx context.Context, req *models.BlogPostCreateRequest) (*models.BlogPost, error) {
	req.LocationID = c.LocationID

	var resp models.BlogPostResponse
	err := c.Do(ctx, "POST", "/blogs/posts", req, &resp)
	if err != nil {
		return nil, fmt.Errorf("create blog post: %w", err)
	}
	return &resp.Post, nil
}

// UpdateBlogPost updates an existing blog post.
func (c *Client) UpdateBlogPost(ctx context.Context, postID string, req *models.BlogPostUpdateRequest) (*models.BlogPost, error) {
	req.LocationID = c.LocationID

	var resp models.BlogPostResponse
	err := c.Do(ctx, "PUT", "/blogs/posts/"+postID, req, &resp)
	if err != nil {
		return nil, fmt.Errorf("update blog post: %w", err)
	}
	return &resp.Post, nil
}

// CheckBlogSlug checks if a blog URL slug exists.
func (c *Client) CheckBlogSlug(ctx context.Context, blogID, slug string) (*models.SlugCheckResponse, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)
	params.Set("blogId", blogID)
	params.Set("slug", slug)

	var resp models.SlugCheckResponse
	err := c.Do(ctx, "GET", "/blogs/posts/url-slug-exists?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("check slug: %w", err)
	}
	return &resp, nil
}

// GetBlogAuthors returns all blog authors for the location.
func (c *Client) GetBlogAuthors(ctx context.Context) ([]models.BlogAuthor, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)

	var resp models.BlogAuthorsResponse
	err := c.Do(ctx, "GET", "/blogs/authors?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get blog authors: %w", err)
	}
	return resp.Authors, nil
}

// GetBlogCategories returns all blog categories for the location.
func (c *Client) GetBlogCategories(ctx context.Context) ([]models.BlogCategory, error) {
	params := url.Values{}
	params.Set("locationId", c.LocationID)

	var resp models.BlogCategoriesResponse
	err := c.Do(ctx, "GET", "/blogs/categories?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("get blog categories: %w", err)
	}
	return resp.Categories, nil
}
