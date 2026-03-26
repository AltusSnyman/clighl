package cmd

import (
	"github.com/spf13/cobra"
)

var socialCmd = &cobra.Command{
	Use:   "social",
	Short: "Manage social media accounts and posts",
	Long:  `View social media accounts, create and manage posts, and check statistics.`,
}

func init() {
	rootCmd.AddCommand(socialCmd)
}
