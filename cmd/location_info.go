package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var locationInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get location details",
	RunE:  runLocationInfo,
}

func init() {
	locationCmd.AddCommand(locationInfoCmd)
}

func runLocationInfo(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	location, err := client.GetLocation(cmd.Context())
	if err != nil {
		return err
	}

	fmt.Print(getFormatter().FormatLocation(location))
	return nil
}
