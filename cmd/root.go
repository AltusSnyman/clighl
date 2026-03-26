package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var (
	version    = "0.1.0"
	jsonOutput bool
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:   "clighl",
	Short: "CLI tool to control Go HighLevel CRM",
	Long: `clighl is a cross-platform CLI tool for managing Go HighLevel CRM.

Control contacts, pipelines, opportunities, and more from your terminal.
All operations are scoped by Location ID and authenticated via access token.

Get started:
  clighl auth          Set up your credentials
  clighl contacts      Manage contacts
  clighl pipelines     View pipelines and stages
  clighl opportunities Manage opportunities`,
	Version: version,
}

func Execute() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Output in JSON format")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose output")
}
