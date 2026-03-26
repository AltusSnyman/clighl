package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blogsAuthorsCmd = &cobra.Command{
	Use:   "authors",
	Short: "List blog authors",
	RunE:  runBlogsAuthors,
}

func init() {
	blogsCmd.AddCommand(blogsAuthorsCmd)
}

func runBlogsAuthors(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	authors, err := client.GetBlogAuthors(cmd.Context())
	if err != nil {
		return err
	}

	if len(authors) == 0 {
		fmt.Println("No blog authors found.")
		return nil
	}

	fmt.Print(getFormatter().FormatBlogAuthors(authors))
	return nil
}
