package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
)

var emailsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an email template",
	RunE:  runEmailsCreate,
}

var (
	emailCreateName    string
	emailCreateSubject string
	emailCreateHTML    string
)

func init() {
	emailsCmd.AddCommand(emailsCreateCmd)
	emailsCreateCmd.Flags().StringVar(&emailCreateName, "name", "", "Template name (required)")
	emailsCreateCmd.Flags().StringVar(&emailCreateSubject, "subject", "", "Email subject")
	emailsCreateCmd.Flags().StringVar(&emailCreateHTML, "html", "", "Email HTML content")
	_ = emailsCreateCmd.MarkFlagRequired("name")
}

func runEmailsCreate(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	req := &models.EmailTemplateCreateRequest{
		Name:    emailCreateName,
		Subject: emailCreateSubject,
		HTML:    emailCreateHTML,
	}

	template, err := client.CreateEmailTemplate(cmd.Context(), req)
	if err != nil {
		return err
	}

	fmt.Println("Email template created successfully!")
	fmt.Println()
	fmt.Print(getFormatter().FormatEmailTemplate(template))
	return nil
}
