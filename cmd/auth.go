package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/api"
	"github.com/altusmusic/clighl/internal/config"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Set up authentication credentials",
	Long:  `Interactively configure your Go HighLevel Location ID and Access Token.`,
	RunE:  runAuth,
}

func init() {
	rootCmd.AddCommand(authCmd)
}

func runAuth(cmd *cobra.Command, args []string) error {
	fmt.Println("Configure Go HighLevel API credentials")
	fmt.Println()

	// Prompt for Location ID
	locPrompt := promptui.Prompt{
		Label:   "Location ID",
		Default: os.Getenv("CLIGHL_LOCATION_ID"),
		Validate: func(input string) error {
			if len(input) == 0 {
				return fmt.Errorf("location ID is required")
			}
			return nil
		},
	}
	locationID, err := locPrompt.Run()
	if err != nil {
		return fmt.Errorf("prompt cancelled: %w", err)
	}

	// Prompt for Access Token
	tokenPrompt := promptui.Prompt{
		Label:   "Access Token (PIT)",
		Default: os.Getenv("CLIGHL_ACCESS_TOKEN"),
		Mask:    '*',
		Validate: func(input string) error {
			if len(input) == 0 {
				return fmt.Errorf("access token is required")
			}
			return nil
		},
	}
	accessToken, err := tokenPrompt.Run()
	if err != nil {
		return fmt.Errorf("prompt cancelled: %w", err)
	}

	// Validate credentials by making a test API call
	fmt.Print("\nValidating credentials... ")
	client := api.NewClientFromToken(locationID, accessToken)
	_, apiErr := client.ListPipelines(cmd.Context())
	if apiErr != nil {
		fmt.Println("FAILED")
		return fmt.Errorf("credentials are invalid: %w", apiErr)
	}
	fmt.Println("OK")

	// Save config
	cfg := &config.Config{
		LocationID:  locationID,
		AccessToken: accessToken,
	}
	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("save config: %w", err)
	}

	configPath, _ := config.ConfigPath()
	fmt.Printf("\nAuthenticated successfully. Config saved to %s\n", configPath)
	return nil
}
