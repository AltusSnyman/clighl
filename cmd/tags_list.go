package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tagsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tags for the location",
	RunE:  runTagsList,
}

func init() {
	tagsCmd.AddCommand(tagsListCmd)
}

func runTagsList(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	tags, err := client.ListTags(cmd.Context())
	if err != nil {
		return err
	}

	if len(tags) == 0 {
		fmt.Println("No tags found.")
		return nil
	}

	fmt.Printf("Found %d tag(s):\n\n", len(tags))
	fmt.Print(getFormatter().FormatTags(tags))
	return nil
}
