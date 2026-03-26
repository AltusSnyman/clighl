package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var tagsRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove tag(s) from a contact",
	Long: `Remove one or more tags from a contact.

Examples:
  clighl tags remove --contact "Dan" --tags "old-lead"
  clighl tags remove --contact "dan@test.com" --tags "inactive,archived"`,
	RunE: runTagsRemove,
}

var (
	tagsRemoveContact string
	tagsRemoveTags    string
)

func init() {
	tagsCmd.AddCommand(tagsRemoveCmd)
	tagsRemoveCmd.Flags().StringVar(&tagsRemoveContact, "contact", "", "Contact name, email, or ID (required)")
	tagsRemoveCmd.Flags().StringVar(&tagsRemoveTags, "tags", "", "Comma-separated tag names (required)")
	tagsRemoveCmd.MarkFlagRequired("contact")
	tagsRemoveCmd.MarkFlagRequired("tags")
}

func runTagsRemove(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)
	contact, err := res.ResolveContact(cmd.Context(), tagsRemoveContact)
	if err != nil {
		return err
	}

	tagList := parseTags(tagsRemoveTags)
	if len(tagList) == 0 {
		return fmt.Errorf("no tags provided")
	}

	err = client.RemoveContactTags(cmd.Context(), contact.ID, tagList)
	if err != nil {
		return err
	}

	fmt.Printf("Removed tag(s) %s from %s.\n", strings.Join(tagList, ", "), contact.DisplayName())
	return nil
}
