package cmd

import (
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Manage tags",
	Long:  `List location tags, and add or remove tags from contacts.`,
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
