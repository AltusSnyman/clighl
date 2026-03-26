package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var socialPostsCmd = &cobra.Command{
	Use:   "posts",
	Short: "List social posts",
	RunE:  runSocialPosts,
}

var (
	socialPostsAccount string
	socialPostsLimit   int
	socialPostsPage    int
)

func init() {
	socialCmd.AddCommand(socialPostsCmd)
	socialPostsCmd.Flags().StringVar(&socialPostsAccount, "account", "", "Filter by account ID")
	socialPostsCmd.Flags().IntVar(&socialPostsLimit, "limit", 20, "Number of results per page")
	socialPostsCmd.Flags().IntVar(&socialPostsPage, "page", 1, "Page number")
}

func runSocialPosts(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	posts, err := client.GetSocialPosts(cmd.Context(), socialPostsAccount, socialPostsLimit, socialPostsPage)
	if err != nil {
		return err
	}

	if len(posts.Posts) == 0 {
		fmt.Println("No social posts found.")
		return nil
	}

	fmt.Print(getFormatter().FormatSocialPosts(posts.Posts))
	return nil
}
