package cmd

import (
	"github.com/spf13/cobra"
)

var blogsCmd = &cobra.Command{
	Use:   "blogs",
	Short: "Manage blog sites and posts",
	Long:  `List blog sites, manage posts, check slugs, and view authors and categories.`,
}

func init() {
	rootCmd.AddCommand(blogsCmd)
}
