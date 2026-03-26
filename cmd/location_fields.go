package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var locationFieldsCmd = &cobra.Command{
	Use:   "fields",
	Short: "List custom field definitions",
	RunE:  runLocationFields,
}

func init() {
	locationCmd.AddCommand(locationFieldsCmd)
}

func runLocationFields(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	fields, err := client.GetCustomFields(cmd.Context())
	if err != nil {
		return err
	}

	if len(fields) == 0 {
		fmt.Println("No custom fields defined.")
		return nil
	}

	fmt.Printf("Custom fields (%d):\n\n", len(fields))
	fmt.Print(getFormatter().FormatCustomFields(fields))
	return nil
}
