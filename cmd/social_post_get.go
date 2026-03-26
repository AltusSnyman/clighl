package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var socialPostGetCmd = &cobra.Command{
	Use:   "get <post-id>",
	Short: "Get a social post by ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runSocialPostGet,
}

func init() {
	socialCmd.AddCommand(socialPostGetCmd)
}

func runSocialPostGet(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	post, err := client.GetSocialPost(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	fmt.Print(getFormatter().FormatSocialPost(post))
	return nil
}
