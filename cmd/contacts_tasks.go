package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var contactsTasksCmd = &cobra.Command{
	Use:   "tasks <contact-name-or-id>",
	Short: "List tasks for a contact",
	Args:  cobra.ExactArgs(1),
	RunE:  runContactsTasks,
}

func init() {
	contactsCmd.AddCommand(contactsTasksCmd)
}

func runContactsTasks(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)
	contact, err := res.ResolveContact(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	tasks, err := client.ListTasks(cmd.Context(), contact.ID)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Printf("No tasks found for %s.\n", contact.DisplayName())
		return nil
	}

	fmt.Printf("Tasks for %s (%d):\n\n", contact.DisplayName(), len(tasks))
	fmt.Print(getFormatter().FormatTasks(tasks))
	return nil
}
