package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/config"
)

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current authentication status",
	RunE:  runAuthStatus,
}

func init() {
	authCmd.AddCommand(authStatusCmd)
}

func runAuthStatus(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Not authenticated. Run `clighl auth` to set up.")
		return nil
	}

	configPath, _ := config.ConfigPath()
	fmt.Printf("Config:      %s\n", configPath)
	fmt.Printf("Location ID: %s\n", cfg.LocationID)
	fmt.Printf("Token:       %s\n", config.MaskToken(cfg.AccessToken))
	fmt.Printf("API Version: %s\n", cfg.APIVersion)
	return nil
}
