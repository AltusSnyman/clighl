package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
	"github.com/altusmusic/clighl/internal/resolver"
)

var opportunitiesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new opportunity",
	RunE:  runOpportunitiesCreate,
}

var (
	oppCreateContact  string
	oppCreatePipeline string
	oppCreateStage    string
	oppCreateName     string
	oppCreateValue    float64
)

func init() {
	opportunitiesCmd.AddCommand(opportunitiesCreateCmd)
	opportunitiesCreateCmd.Flags().StringVar(&oppCreateContact, "contact", "", "Contact name or ID (required)")
	opportunitiesCreateCmd.Flags().StringVar(&oppCreatePipeline, "pipeline", "", "Pipeline name (required)")
	opportunitiesCreateCmd.Flags().StringVar(&oppCreateStage, "stage", "", "Stage name (required)")
	opportunitiesCreateCmd.Flags().StringVar(&oppCreateName, "name", "", "Opportunity name (defaults to contact name)")
	opportunitiesCreateCmd.Flags().Float64Var(&oppCreateValue, "value", 0, "Monetary value")
	opportunitiesCreateCmd.MarkFlagRequired("contact")
	opportunitiesCreateCmd.MarkFlagRequired("pipeline")
	opportunitiesCreateCmd.MarkFlagRequired("stage")
}

func runOpportunitiesCreate(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)

	// Resolve contact
	contact, err := res.ResolveContact(cmd.Context(), oppCreateContact)
	if err != nil {
		return err
	}

	// Resolve pipeline and stage
	pipeline, err := res.ResolvePipeline(cmd.Context(), oppCreatePipeline)
	if err != nil {
		return err
	}

	stage, err := res.ResolveStage(pipeline, oppCreateStage)
	if err != nil {
		return err
	}

	name := oppCreateName
	if name == "" {
		name = contact.DisplayName()
	}

	req := &models.OpportunityCreateRequest{
		PipelineID:      pipeline.ID,
		PipelineStageID: stage.ID,
		ContactID:       contact.ID,
		Name:            name,
		MonetaryValue:   oppCreateValue,
	}

	opp, err := client.CreateOpportunity(cmd.Context(), req)
	if err != nil {
		return err
	}

	fmt.Printf("Opportunity created: %s → %s > %s\n\n", contact.DisplayName(), pipeline.Name, stage.Name)
	fmt.Print(getFormatter().FormatOpportunity(opp))
	return nil
}
