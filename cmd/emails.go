package cmd

import (
	"github.com/spf13/cobra"
)

var emailsCmd = &cobra.Command{
	Use:   "emails",
	Short: "Manage email templates",
	Long:  `List, create, and manage email templates.`,
}

func init() {
	rootCmd.AddCommand(emailsCmd)
}
