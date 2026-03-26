package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
)

var contactsUpsertCmd = &cobra.Command{
	Use:   "upsert",
	Short: "Create or update a contact (matched by email/phone)",
	Long: `Upsert a contact — creates if not found, updates if matched by email or phone.

Examples:
  clighl contacts upsert --email "dan@test.com" --first-name "Dan" --last-name "Smith"
  clighl contacts upsert --phone "+1234567890" --first-name "Jane" --company "Acme"`,
	RunE: runContactsUpsert,
}

var (
	upsertFirstName string
	upsertLastName  string
	upsertEmail     string
	upsertPhone     string
	upsertCompany   string
	upsertSource    string
)

func init() {
	contactsCmd.AddCommand(contactsUpsertCmd)
	contactsUpsertCmd.Flags().StringVar(&upsertFirstName, "first-name", "", "First name")
	contactsUpsertCmd.Flags().StringVar(&upsertLastName, "last-name", "", "Last name")
	contactsUpsertCmd.Flags().StringVar(&upsertEmail, "email", "", "Email address")
	contactsUpsertCmd.Flags().StringVar(&upsertPhone, "phone", "", "Phone number")
	contactsUpsertCmd.Flags().StringVar(&upsertCompany, "company", "", "Company name")
	contactsUpsertCmd.Flags().StringVar(&upsertSource, "source", "clighl", "Contact source")
}

func runContactsUpsert(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	if upsertEmail == "" && upsertPhone == "" {
		return fmt.Errorf("at least --email or --phone is required for upsert matching")
	}

	req := &models.ContactUpsertRequest{
		FirstName:   upsertFirstName,
		LastName:    upsertLastName,
		Email:       upsertEmail,
		Phone:       upsertPhone,
		CompanyName: upsertCompany,
		Source:      upsertSource,
	}

	contact, err := client.UpsertContact(cmd.Context(), req)
	if err != nil {
		return err
	}

	fmt.Printf("Contact upserted: %s\n\n", contact.DisplayName())
	fmt.Print(getFormatter().FormatContact(contact))
	return nil
}
