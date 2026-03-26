package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blogsPostsCmd = &cobra.Command{
	Use:   "posts <blog-id>",
	Short: "List blog posts",
	Args:  cobra.ExactArgs(1),
	RunE:  runBlogsPosts,
}

var (
	blogPostsLimit  int
	blogPostsOffset int
)

func init() {
	blogsCmd.AddCommand(blogsPostsCmd)
	blogsPostsCmd.Flags().IntVar(&blogPostsLimit, "limit", 20, "Number of results to return")
	blogsPostsCmd.Flags().IntVar(&blogPostsOffset, "offset", 0, "Offset for pagination")
}

func runBlogsPosts(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	posts, err := client.GetBlogPosts(cmd.Context(), args[0], blogPostsLimit, blogPostsOffset)
	if err != nil {
		return err
	}

	if len(posts.Posts) == 0 {
		fmt.Println("No blog posts found.")
		return nil
	}

	fmt.Print(getFormatter().FormatBlogPosts(posts.Posts))
	return nil
}
