package cmd

import (
	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/api"
	"github.com/altusmusic/clighl/internal/config"
	"github.com/altusmusic/clighl/internal/output"
)

var contactsCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Manage contacts",
	Long:  `Search, view, create, and manage Go HighLevel contacts.`,
}

func init() {
	rootCmd.AddCommand(contactsCmd)
}

func newAPIClient() (*api.Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	return api.NewClient(cfg), nil
}

func getFormatter() output.Formatter {
	if jsonOutput {
		return &output.JSONFormatter{}
	}
	return &output.TableFormatter{}
}
