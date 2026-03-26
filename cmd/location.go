package cmd

import (
	"github.com/spf13/cobra"
)

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "View location/account information",
	Long:  `View details and settings for the current GHL location/sub-account.`,
}

func init() {
	rootCmd.AddCommand(locationCmd)
}
