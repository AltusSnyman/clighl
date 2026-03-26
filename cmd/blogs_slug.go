package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blogsSlugCheckCmd = &cobra.Command{
	Use:   "slug-check",
	Short: "Check if a URL slug is available",
	RunE:  runBlogsSlugCheck,
}

var (
	slugCheckBlog string
	slugCheckSlug string
)

func init() {
	blogsCmd.AddCommand(blogsSlugCheckCmd)
	blogsSlugCheckCmd.Flags().StringVar(&slugCheckBlog, "blog", "", "Blog ID (required)")
	blogsSlugCheckCmd.Flags().StringVar(&slugCheckSlug, "slug", "", "URL slug to check (required)")
	_ = blogsSlugCheckCmd.MarkFlagRequired("blog")
	_ = blogsSlugCheckCmd.MarkFlagRequired("slug")
}

func runBlogsSlugCheck(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	resp, err := client.CheckBlogSlug(cmd.Context(), slugCheckBlog, slugCheckSlug)
	if err != nil {
		return err
	}

	if resp.Exists {
		fmt.Printf("Slug '%s' already exists.\n", slugCheckSlug)
	} else {
		fmt.Printf("Slug '%s' is available.\n", slugCheckSlug)
	}
	return nil
}
