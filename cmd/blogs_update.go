package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
)

var blogsUpdateCmd = &cobra.Command{
	Use:   "update <post-id>",
	Short: "Update a blog post",
	Args:  cobra.ExactArgs(1),
	RunE:  runBlogsUpdate,
}

var (
	blogUpdateBlog     string
	blogUpdateTitle    string
	blogUpdateHTML     string
	blogUpdateStatus   string
	blogUpdateAuthor   string
	blogUpdateCategory string
	blogUpdateSlug     string
)

func init() {
	blogsCmd.AddCommand(blogsUpdateCmd)
	blogsUpdateCmd.Flags().StringVar(&blogUpdateBlog, "blog", "", "Blog ID (required)")
	blogsUpdateCmd.Flags().StringVar(&blogUpdateTitle, "title", "", "Post title")
	blogsUpdateCmd.Flags().StringVar(&blogUpdateHTML, "html", "", "Post HTML content")
	blogsUpdateCmd.Flags().StringVar(&blogUpdateStatus, "status", "", "Post status")
	blogsUpdateCmd.Flags().StringVar(&blogUpdateAuthor, "author", "", "Author ID")
	blogsUpdateCmd.Flags().StringVar(&blogUpdateCategory, "category", "", "Category ID")
	blogsUpdateCmd.Flags().StringVar(&blogUpdateSlug, "slug", "", "URL slug")
	_ = blogsUpdateCmd.MarkFlagRequired("blog")
}

func runBlogsUpdate(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	req := &models.BlogPostUpdateRequest{}
	hasUpdate := false

	if cmd.Flags().Changed("title") {
		req.Title = blogUpdateTitle
		hasUpdate = true
	}
	if cmd.Flags().Changed("html") {
		req.RawHTML = blogUpdateHTML
		hasUpdate = true
	}
	if cmd.Flags().Changed("status") {
		req.Status = blogUpdateStatus
		hasUpdate = true
	}
	if cmd.Flags().Changed("author") {
		req.Author = blogUpdateAuthor
		hasUpdate = true
	}
	if cmd.Flags().Changed("category") {
		req.CategoryID = blogUpdateCategory
		hasUpdate = true
	}
	if cmd.Flags().Changed("slug") {
		req.Slug = blogUpdateSlug
		hasUpdate = true
	}

	if !hasUpdate {
		return fmt.Errorf("no fields to update. Use flags like --title, --html, --status, etc.")
	}

	post, err := client.UpdateBlogPost(cmd.Context(), args[0], req)
	if err != nil {
		return err
	}

	fmt.Println("Blog post updated successfully!")
	fmt.Println()
	fmt.Print(getFormatter().FormatBlogPost(post))
	return nil
}
