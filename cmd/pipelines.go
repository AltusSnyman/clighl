package cmd

import (
	"github.com/spf13/cobra"
)

var pipelinesCmd = &cobra.Command{
	Use:   "pipelines",
	Short: "View pipelines and stages",
	Long:  `List and inspect Go HighLevel pipelines and their stages.`,
}

func init() {
	rootCmd.AddCommand(pipelinesCmd)
}
