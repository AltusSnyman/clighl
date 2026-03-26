package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blogsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List blogs",
	RunE:  runBlogsList,
}

func init() {
	blogsCmd.AddCommand(blogsListCmd)
}

func runBlogsList(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	blogs, err := client.ListBlogs(cmd.Context())
	if err != nil {
		return err
	}

	if len(blogs) == 0 {
		fmt.Println("No blogs found.")
		return nil
	}

	fmt.Print(getFormatter().FormatBlogs(blogs))
	return nil
}
