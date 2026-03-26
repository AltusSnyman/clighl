package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
	"github.com/altusmusic/clighl/internal/resolver"
)

var contactsUpdateCmd = &cobra.Command{
	Use:   "update <contact-name-or-id>",
	Short: "Update a contact",
	Long: `Update fields on an existing contact.

Examples:
  clighl contacts update "Dan" --email "new@email.com"
  clighl contacts update "Dan" --first-name "Daniel" --company "Acme Inc"`,
	Args: cobra.ExactArgs(1),
	RunE: runContactsUpdate,
}

var (
	updateFirstName string
	updateLastName  string
	updateEmail     string
	updatePhone     string
	updateCompany   string
)

func init() {
	contactsCmd.AddCommand(contactsUpdateCmd)
	contactsUpdateCmd.Flags().StringVar(&updateFirstName, "first-name", "", "First name")
	contactsUpdateCmd.Flags().StringVar(&updateLastName, "last-name", "", "Last name")
	contactsUpdateCmd.Flags().StringVar(&updateEmail, "email", "", "Email address")
	contactsUpdateCmd.Flags().StringVar(&updatePhone, "phone", "", "Phone number")
	contactsUpdateCmd.Flags().StringVar(&updateCompany, "company", "", "Company name")
}

func runContactsUpdate(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)
	contact, err := res.ResolveContact(cmd.Context(), args[0])
	if err != nil {
		return err
	}

	req := &models.ContactUpdateRequest{}
	hasUpdate := false

	if cmd.Flags().Changed("first-name") {
		req.FirstName = updateFirstName
		hasUpdate = true
	}
	if cmd.Flags().Changed("last-name") {
		req.LastName = updateLastName
		hasUpdate = true
	}
	if cmd.Flags().Changed("email") {
		req.Email = updateEmail
		hasUpdate = true
	}
	if cmd.Flags().Changed("phone") {
		req.Phone = updatePhone
		hasUpdate = true
	}
	if cmd.Flags().Changed("company") {
		req.CompanyName = updateCompany
		hasUpdate = true
	}

	if !hasUpdate {
		return fmt.Errorf("no fields to update. Use flags like --email, --phone, --first-name, etc.")
	}

	updated, err := client.UpdateContact(cmd.Context(), contact.ID, req)
	if err != nil {
		return err
	}

	fmt.Printf("Contact %s updated.\n\n", contact.DisplayName())
	fmt.Print(getFormatter().FormatContact(updated))
	return nil
}
