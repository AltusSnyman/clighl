package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
)

var blogsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a blog post",
	RunE:  runBlogsCreate,
}

var (
	blogCreateBlog     string
	blogCreateTitle    string
	blogCreateHTML     string
	blogCreateStatus   string
	blogCreateAuthor   string
	blogCreateCategory string
	blogCreateSlug     string
)

func init() {
	blogsCmd.AddCommand(blogsCreateCmd)
	blogsCreateCmd.Flags().StringVar(&blogCreateBlog, "blog", "", "Blog ID (required)")
	blogsCreateCmd.Flags().StringVar(&blogCreateTitle, "title", "", "Post title (required)")
	blogsCreateCmd.Flags().StringVar(&blogCreateHTML, "html", "", "Post HTML content")
	blogsCreateCmd.Flags().StringVar(&blogCreateStatus, "status", "draft", "Post status")
	blogsCreateCmd.Flags().StringVar(&blogCreateAuthor, "author", "", "Author ID")
	blogsCreateCmd.Flags().StringVar(&blogCreateCategory, "category", "", "Category ID")
	blogsCreateCmd.Flags().StringVar(&blogCreateSlug, "slug", "", "URL slug")
	_ = blogsCreateCmd.MarkFlagRequired("blog")
	_ = blogsCreateCmd.MarkFlagRequired("title")
}

func runBlogsCreate(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	req := &models.BlogPostCreateRequest{
		BlogID:     blogCreateBlog,
		Title:      blogCreateTitle,
		RawHTML:    blogCreateHTML,
		Status:     blogCreateStatus,
		Author:     blogCreateAuthor,
		CategoryID: blogCreateCategory,
		Slug:       blogCreateSlug,
	}

	post, err := client.CreateBlogPost(cmd.Context(), req)
	if err != nil {
		return err
	}

	fmt.Println("Blog post created successfully!")
	fmt.Println()
	fmt.Print(getFormatter().FormatBlogPost(post))
	return nil
}
