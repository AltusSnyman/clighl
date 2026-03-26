package cmd

import (
	"github.com/spf13/cobra"
)

var opportunitiesCmd = &cobra.Command{
	Use:     "opportunities",
	Aliases: []string{"opp"},
	Short:   "Manage opportunities",
	Long:    `List, create, and move opportunities between pipeline stages.`,
}

func init() {
	rootCmd.AddCommand(opportunitiesCmd)
}
