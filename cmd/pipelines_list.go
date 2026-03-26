package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pipelinesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all pipelines",
	RunE:  runPipelinesList,
}

func init() {
	pipelinesCmd.AddCommand(pipelinesListCmd)
}

func runPipelinesList(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	pipelines, err := client.ListPipelines(cmd.Context())
	if err != nil {
		return err
	}

	if len(pipelines) == 0 {
		fmt.Println("No pipelines found.")
		return nil
	}

	fmt.Printf("Found %d pipeline(s):\n\n", len(pipelines))
	fmt.Print(getFormatter().FormatPipelines(pipelines))
	return nil
}
