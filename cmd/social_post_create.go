package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
)

var socialPostCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a social post",
	RunE:  runSocialPostCreate,
}

var (
	socialCreateAccounts string
	socialCreateContent  string
	socialCreateSummary  string
	socialCreateSchedule string
)

func init() {
	socialCmd.AddCommand(socialPostCreateCmd)
	socialPostCreateCmd.Flags().StringVar(&socialCreateAccounts, "accounts", "", "Comma-separated account IDs (required)")
	socialPostCreateCmd.Flags().StringVar(&socialCreateContent, "content", "", "Post content (required)")
	socialPostCreateCmd.Flags().StringVar(&socialCreateSummary, "summary", "", "Post summary")
	socialPostCreateCmd.Flags().StringVar(&socialCreateSchedule, "schedule", "", "Schedule time (ISO 8601)")
	_ = socialPostCreateCmd.MarkFlagRequired("accounts")
	_ = socialPostCreateCmd.MarkFlagRequired("content")
}

func runSocialPostCreate(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	accountIDs := strings.Split(socialCreateAccounts, ",")

	req := &models.SocialPostCreateRequest{
		AccountIDs: accountIDs,
		Content:    socialCreateContent,
		Summary:    socialCreateSummary,
		ScheduledAt: socialCreateSchedule,
	}

	post, err := client.CreateSocialPost(cmd.Context(), req)
	if err != nil {
		return err
	}

	fmt.Println("Social post created successfully!")
	fmt.Println()
	fmt.Print(getFormatter().FormatSocialPost(post))
	return nil
}
