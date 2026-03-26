package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/config"
)

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored credentials",
	RunE:  runAuthLogout,
}

func init() {
	authCmd.AddCommand(authLogoutCmd)
}

func runAuthLogout(cmd *cobra.Command, args []string) error {
	if err := config.Delete(); err != nil {
		return err
	}
	fmt.Println("Logged out. Config removed.")
	return nil
}
