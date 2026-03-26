package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/resolver"
)

var tagsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tag(s) to a contact",
	Long: `Add one or more tags to a contact.

Examples:
  clighl tags add --contact "Dan" --tags "vip"
  clighl tags add --contact "dan@test.com" --tags "lead,hot,priority"`,
	RunE: runTagsAdd,
}

var (
	tagsAddContact string
	tagsAddTags    string
)

func init() {
	tagsCmd.AddCommand(tagsAddCmd)
	tagsAddCmd.Flags().StringVar(&tagsAddContact, "contact", "", "Contact name, email, or ID (required)")
	tagsAddCmd.Flags().StringVar(&tagsAddTags, "tags", "", "Comma-separated tag names (required)")
	tagsAddCmd.MarkFlagRequired("contact")
	tagsAddCmd.MarkFlagRequired("tags")
}

func runTagsAdd(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)
	contact, err := res.ResolveContact(cmd.Context(), tagsAddContact)
	if err != nil {
		return err
	}

	tagList := parseTags(tagsAddTags)
	if len(tagList) == 0 {
		return fmt.Errorf("no tags provided")
	}

	resultTags, err := client.AddContactTags(cmd.Context(), contact.ID, tagList)
	if err != nil {
		return err
	}

	fmt.Printf("Added tag(s) %s to %s.\n", strings.Join(tagList, ", "), contact.DisplayName())
	if len(resultTags) > 0 {
		fmt.Printf("Contact now has tags: %s\n", strings.Join(resultTags, ", "))
	}
	return nil
}

func parseTags(input string) []string {
	parts := strings.Split(input, ",")
	tags := make([]string, 0, len(parts))
	for _, p := range parts {
		t := strings.TrimSpace(p)
		if t != "" {
			tags = append(tags, t)
		}
	}
	return tags
}
