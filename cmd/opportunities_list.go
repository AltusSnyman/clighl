package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var opportunitiesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List opportunities",
	RunE:  runOpportunitiesList,
}

var (
	oppListPipeline string
	oppListLimit    int
	oppListPage     int
)

func init() {
	opportunitiesCmd.AddCommand(opportunitiesListCmd)
	opportunitiesListCmd.Flags().StringVar(&oppListPipeline, "pipeline", "", "Filter by pipeline ID")
	opportunitiesListCmd.Flags().IntVar(&oppListLimit, "limit", 20, "Number of results per page")
	opportunitiesListCmd.Flags().IntVar(&oppListPage, "page", 1, "Page number")
}

func runOpportunitiesList(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	resp, err := client.SearchOpportunities(cmd.Context(), oppListPipeline, "", oppListPage, oppListLimit)
	if err != nil {
		return err
	}

	if len(resp.Opportunities) == 0 {
		fmt.Println("No opportunities found.")
		return nil
	}

	fmt.Printf("Found %d opportunity(ies):\n\n", len(resp.Opportunities))
	fmt.Print(getFormatter().FormatOpportunities(resp.Opportunities))
	return nil
}
