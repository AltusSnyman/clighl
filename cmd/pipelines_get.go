package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pipelinesGetCmd = &cobra.Command{
	Use:   "get <id-or-name>",
	Short: "Get pipeline details",
	Args:  cobra.ExactArgs(1),
	RunE:  runPipelinesGet,
}

func init() {
	pipelinesCmd.AddCommand(pipelinesGetCmd)
}

func runPipelinesGet(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	// List all pipelines and find by ID or name
	pipelines, err := client.ListPipelines(cmd.Context())
	if err != nil {
		return err
	}

	query := args[0]
	for i := range pipelines {
		if pipelines[i].ID == query || pipelines[i].Name == query {
			fmt.Print(getFormatter().FormatPipeline(&pipelines[i]))
			return nil
		}
	}

	return fmt.Errorf("pipeline '%s' not found", query)
}
