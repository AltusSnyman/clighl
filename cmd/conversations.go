package cmd

import (
	"github.com/spf13/cobra"
)

var conversationsCmd = &cobra.Command{
	Use:     "conversations",
	Aliases: []string{"conv"},
	Short:   "Manage conversations and messages",
	Long:    `Search conversations, view messages, and send messages to contacts.`,
}

func init() {
	rootCmd.AddCommand(conversationsCmd)
}
