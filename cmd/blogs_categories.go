package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blogsCategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "List blog categories",
	RunE:  runBlogsCategories,
}

func init() {
	blogsCmd.AddCommand(blogsCategoriesCmd)
}

func runBlogsCategories(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	categories, err := client.GetBlogCategories(cmd.Context())
	if err != nil {
		return err
	}

	if len(categories) == 0 {
		fmt.Println("No blog categories found.")
		return nil
	}

	fmt.Print(getFormatter().FormatBlogCategories(categories))
	return nil
}
