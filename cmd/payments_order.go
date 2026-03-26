package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var paymentsOrderCmd = &cobra.Command{
	Use:   "order <id>",
	Short: "Get order by ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runPaymentsOrder,
}

func init() {
	paymentsCmd.AddCommand(paymentsOrderCmd)
}

func runPaymentsOrder(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	order, err := client.GetOrder(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	fmt.Print(getFormatter().FormatOrder(order))
	return nil
}
