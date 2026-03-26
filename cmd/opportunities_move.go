package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/altusmusic/clighl/internal/models"
	"github.com/altusmusic/clighl/internal/resolver"
)

var opportunitiesMoveCmd = &cobra.Command{
	Use:   "move <contact-name>",
	Short: "Move a contact to a pipeline stage",
	Long: `Move a contact to a specific pipeline and stage.

If the contact already has an opportunity in the pipeline, it updates the stage.
If not, it creates a new opportunity.

Examples:
  clighl opportunities move "Altus" --pipeline "Sales" --stage "Lead"
  clighl opportunities move "john@example.com" --pipeline "Onboarding" --stage "Qualified"`,
	Args: cobra.ExactArgs(1),
	RunE: runOpportunitiesMove,
}

var (
	movePipeline string
	moveStage    string
	moveValue    float64
)

func init() {
	opportunitiesCmd.AddCommand(opportunitiesMoveCmd)
	opportunitiesMoveCmd.Flags().StringVar(&movePipeline, "pipeline", "", "Pipeline name (required)")
	opportunitiesMoveCmd.Flags().StringVar(&moveStage, "stage", "", "Stage name (required)")
	opportunitiesMoveCmd.Flags().Float64Var(&moveValue, "value", 0, "Monetary value")
	opportunitiesMoveCmd.MarkFlagRequired("pipeline")
	opportunitiesMoveCmd.MarkFlagRequired("stage")
}

func runOpportunitiesMove(cmd *cobra.Command, args []string) error {
	client, err := newAPIClient()
	if err != nil {
		return err
	}

	res := resolver.NewResolver(client)

	// 1. Resolve contact
	fmt.Printf("Searching for contact '%s'...\n", args[0])
	contact, err := res.ResolveContact(cmd.Context(), args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Found: %s (%s)\n", contact.DisplayName(), contact.ID)

	// 2. Resolve pipeline
	pipeline, err := res.ResolvePipeline(cmd.Context(), movePipeline)
	if err != nil {
		return err
	}

	// 3. Resolve stage
	stage, err := res.ResolveStage(pipeline, moveStage)
	if err != nil {
		return err
	}

	// 4. Check for existing opportunity in this pipeline
	existing, err := client.SearchOpportunities(cmd.Context(), pipeline.ID, contact.ID, 1, 10)
	if err != nil {
		return fmt.Errorf("search existing opportunities: %w", err)
	}

	if len(existing.Opportunities) > 0 {
		// Update existing opportunity
		opp := existing.Opportunities[0]
		fmt.Printf("Updating existing opportunity '%s'...\n", opp.Name)

		updateReq := &models.OpportunityUpdateRequest{
			PipelineStageID: stage.ID,
		}
		if moveValue > 0 {
			updateReq.MonetaryValue = moveValue
		}

		updated, err := client.UpdateOpportunity(cmd.Context(), opp.ID, updateReq)
		if err != nil {
			return err
		}

		fmt.Printf("\nMoved %s to %s > %s\n", contact.DisplayName(), pipeline.Name, stage.Name)
		if jsonOutput {
			fmt.Print(getFormatter().FormatOpportunity(updated))
		}
	} else {
		// Create new opportunity
		fmt.Printf("Creating new opportunity...\n")

		createReq := &models.OpportunityCreateRequest{
			PipelineID:      pipeline.ID,
			PipelineStageID: stage.ID,
			ContactID:       contact.ID,
			Name:            contact.DisplayName(),
			MonetaryValue:   moveValue,
		}

		created, err := client.CreateOpportunity(cmd.Context(), createReq)
		if err != nil {
			return err
		}

		fmt.Printf("\nMoved %s to %s > %s (new opportunity)\n", contact.DisplayName(), pipeline.Name, stage.Name)
		if jsonOutput {
			fmt.Print(getFormatter().FormatOpportunity(created))
		}
	}

	return nil
}
