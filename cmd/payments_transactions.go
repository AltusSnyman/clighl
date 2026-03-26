package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var paymentsTransactionsCmd = &cobra.Command{
	Use:   "transactions",
	Short: "List transactions",
	RunE:  runPaymentsTransactions,
}

var (
	txnContact string
	txnLimit   int
	txnPage    int
)

func init() {
	paymentsCmd.AddCommand(paymentsTransactionsCmd)
	paymentsTransactionsCmd.Flags().StringVar(&txnContact, "contact", "", "Filter by contact name or ID")
	paymentsTransactionsCmd.Flags().IntVar(&txnLimit, "limit", 20, "Number of results per page")
	paymentsTransactionsCmd.Flags().IntVar(&txnPage, "page", 1, "Page number")
}

func runPaymentsTransactions(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	contactID := ""
	if txnContact != "" {
		res := resolver.NewResolver(client)
		contact, err := res.ResolveContact(cmd.Context(), txnContact)
		if err != nil {
			return err
		}
		contactID = contact.ID
	}

	resp, err := client.ListTransactions(cmd.Context(), contactID, txnLimit, txnPage)
	if err != nil {
		return err
	}

	if len(resp.Transactions) == 0 {
		fmt.Println("No transactions found.")
		return nil
	}

	fmt.Printf("Transactions (page %d, %d total):\n\n", txnPage, resp.Total)
	fmt.Print(getFormatter().FormatTransactions(resp.Transactions))
	return nil
}
