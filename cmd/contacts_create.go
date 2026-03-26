package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
)

var contactsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new contact",
	Long:  `Create a new contact. Use flags for non-interactive mode, or omit for interactive prompts.`,
	RunE:  runContactsCreate,
}

var (
	createFirstName string
	createLastName  string
	createEmail     string
	createPhone     string
	createCompany   string
)

func init() {
	contactsCmd.AddCommand(contactsCreateCmd)
	contactsCreateCmd.Flags().StringVar(&createFirstName, "first-name", "", "First name")
	contactsCreateCmd.Flags().StringVar(&createLastName, "last-name", "", "Last name")
	contactsCreateCmd.Flags().StringVar(&createEmail, "email", "", "Email address")
	contactsCreateCmd.Flags().StringVar(&createPhone, "phone", "", "Phone number")
	contactsCreateCmd.Flags().StringVar(&createCompany, "company", "", "Company name")
}

func runContactsCreate(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	// If no flags provided, use interactive mode
	if createFirstName == "" && createLastName == "" && createEmail == "" && createPhone == "" {
		if err := promptContactFields(); err != nil {
			return err
		}
	}

	req := &models.ContactCreateRequest{
		FirstName:   createFirstName,
		LastName:    createLastName,
		Email:       createEmail,
		Phone:       createPhone,
		CompanyName: createCompany,
		Source:      "clighl",
	}

	contact, err := client.CreateContact(cmd.Context(), req)
	if err != nil {
		return err
	}

	fmt.Println("Contact created successfully!")
	fmt.Println()
	fmt.Print(getFormatter().FormatContact(contact))
	return nil
}

func promptContactFields() error {
	prompt := func(label, defaultVal string) (string, error) {
		p := promptui.Prompt{Label: label, Default: defaultVal}
		return p.Run()
	}

	var err error
	createFirstName, err = prompt("First name", "")
	if err != nil {
		return err
	}
	createLastName, err = prompt("Last name", "")
	if err != nil {
		return err
	}
	createEmail, err = prompt("Email", "")
	if err != nil {
		return err
	}
	createPhone, err = prompt("Phone", "")
	if err != nil {
		return err
	}
	createCompany, err = prompt("Company", "")
	if err != nil {
		return err
	}
	return nil
}
