package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var opportunitiesGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get opportunity details by ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runOpportunitiesGet,
}

func init() {
	opportunitiesCmd.AddCommand(opportunitiesGetCmd)
}

func runOpportunitiesGet(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	opp, err := client.GetOpportunity(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	fmt.Print(getFormatter().FormatOpportunity(opp))
	return nil
}
