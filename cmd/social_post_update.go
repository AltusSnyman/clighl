package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
)

var socialPostUpdateCmd = &cobra.Command{
	Use:   "update <post-id>",
	Short: "Update a social post",
	Args:  cobra.ExactArgs(1),
	RunE:  runSocialPostUpdate,
}

var (
	socialUpdateContent  string
	socialUpdateSummary  string
	socialUpdateSchedule string
)

func init() {
	socialCmd.AddCommand(socialPostUpdateCmd)
	socialPostUpdateCmd.Flags().StringVar(&socialUpdateContent, "content", "", "Post content")
	socialPostUpdateCmd.Flags().StringVar(&socialUpdateSummary, "summary", "", "Post summary")
	socialPostUpdateCmd.Flags().StringVar(&socialUpdateSchedule, "schedule", "", "Schedule time (ISO 8601)")
}

func runSocialPostUpdate(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	req := &models.SocialPostUpdateRequest{}
	hasUpdate := false

	if cmd.Flags().Changed("content") {
		req.Content = socialUpdateContent
		hasUpdate = true
	}
	if cmd.Flags().Changed("summary") {
		req.Summary = socialUpdateSummary
		hasUpdate = true
	}
	if cmd.Flags().Changed("schedule") {
		req.ScheduledAt = socialUpdateSchedule
		hasUpdate = true
	}

	if !hasUpdate {
		return fmt.Errorf("no fields to update. Use flags like --content, --summary, --schedule")
	}

	post, err := client.UpdateSocialPost(cmd.Context(), args[0], req)
	if err != nil {
		return err
	}

	fmt.Println("Social post updated successfully!")
	fmt.Println()
	fmt.Print(getFormatter().FormatSocialPost(post))
	return nil
}
