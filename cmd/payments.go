package cmd

import (
	"github.com/spf13/cobra"
)

var paymentsCmd = &cobra.Command{
	Use:     "payments",
	Aliases: []string{"pay"},
	Short:   "View payment orders and transactions",
	Long:    `View orders by ID and list transactions with optional contact filtering.`,
}

func init() {
	rootCmd.AddCommand(paymentsCmd)
}
