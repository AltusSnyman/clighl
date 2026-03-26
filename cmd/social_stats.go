package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var socialStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get social media statistics",
	RunE:  runSocialStats,
}

var socialStatsAccount string

func init() {
	socialCmd.AddCommand(socialStatsCmd)
	socialStatsCmd.Flags().StringVar(&socialStatsAccount, "account", "", "Comma-separated account IDs (optional)")
}

func runSocialStats(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	var accountIDs []string
	if socialStatsAccount != "" {
		accountIDs = strings.Split(socialStatsAccount, ",")
	}

	stats, err := client.GetSocialStats(cmd.Context(), accountIDs)
	if err != nil {
		return err
	}

	if jsonOutput {
		data, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	}

	for _, s := range stats.Stats {
		fmt.Printf("Account:    %s\n", s.AccountID)
		fmt.Printf("Platform:   %s\n", s.Platform)
		fmt.Printf("Followers:  %v\n", s.Followers)
		fmt.Printf("Posts:      %v\n", s.Posts)
		fmt.Printf("Engagement: %v\n", s.Engagement)
		fmt.Println()
	}
	return nil
}
