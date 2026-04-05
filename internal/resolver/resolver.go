package resolver

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"golang.org/x/term"

	"github.com/altusmusic/clighl/internal/api"
	"github.com/altusmusic/clighl/internal/models"
)

// Resolver translates human-readable names into GHL API IDs.
type Resolver struct {
	Client *api.Client
}

// NewResolver creates a new resolver.
func NewResolver(client *api.Client) *Resolver {
	return &Resolver{Client: client}
}

// ResolveContact searches for a contact by name/email and handles disambiguation.
func (r *Resolver) ResolveContact(ctx context.Context, query string) (*models.Contact, error) {
	// If the caller already provided a contact ID, fetch it directly.
	if looksLikeContactID(query) {
		contact, err := r.Client.GetContact(ctx, query)
		if err == nil && contact != nil {
			return contact, nil
		}
	}

	resp, err := r.Client.SearchContacts(ctx, query, 1, 20)
	if err != nil {
		return nil, fmt.Errorf("search contacts: %w", err)
	}

	if len(resp.Contacts) == 0 {
		return nil, fmt.Errorf("no contact found matching '%s'", query)
	}

	if len(resp.Contacts) == 1 {
		return &resp.Contacts[0], nil
	}

	// Prefer exact email/name matches in non-interactive flows.
	for i := range resp.Contacts {
		c := &resp.Contacts[i]
		if strings.EqualFold(c.Email, query) || strings.EqualFold(c.DisplayName(), query) || strings.EqualFold(c.ID, query) {
			return c, nil
		}
	}

	// Multiple matches — need disambiguation
	if !isTerminal() {
		return nil, fmt.Errorf("multiple contacts match '%s'. Use a unique email or exact contact ID", query)
	}

	return disambiguateContact(resp.Contacts)
}

func looksLikeContactID(query string) bool {
	if len(query) < 10 || strings.Contains(query, "@") || strings.Contains(query, " ") {
		return false
	}
	for _, r := range query {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') && r != '-' && r != '_' {
			return false
		}
	}
	return true
}

// ResolvePipeline finds a pipeline by name (case-insensitive).
func (r *Resolver) ResolvePipeline(ctx context.Context, name string) (*models.Pipeline, error) {
	pipelines, err := r.Client.ListPipelines(ctx)
	if err != nil {
		return nil, fmt.Errorf("list pipelines: %w", err)
	}

	var match *models.Pipeline
	var available []string
	for i := range pipelines {
		available = append(available, pipelines[i].Name)
		if strings.EqualFold(pipelines[i].Name, name) {
			match = &pipelines[i]
		}
	}

	if match == nil {
		return nil, fmt.Errorf("no pipeline found matching '%s'. Available: %s", name, strings.Join(available, ", "))
	}

	return match, nil
}

// ResolveStage finds a stage within a pipeline by name (case-insensitive).
func (r *Resolver) ResolveStage(pipeline *models.Pipeline, stageName string) (*models.Stage, error) {
	var available []string
	for i := range pipeline.Stages {
		available = append(available, pipeline.Stages[i].Name)
		if strings.EqualFold(pipeline.Stages[i].Name, stageName) {
			return &pipeline.Stages[i], nil
		}
	}

	return nil, fmt.Errorf("no stage '%s' found in pipeline '%s'. Available: %s",
		stageName, pipeline.Name, strings.Join(available, ", "))
}

func isTerminal() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}

// ResolveCalendar finds a calendar by name (case-insensitive).
func (r *Resolver) ResolveCalendar(ctx context.Context, name string) (*models.Calendar, error) {
	calendars, err := r.Client.ListCalendars(ctx)
	if err != nil {
		return nil, fmt.Errorf("list calendars: %w", err)
	}

	var match *models.Calendar
	var available []string
	for i := range calendars {
		available = append(available, calendars[i].Name)
		if strings.EqualFold(calendars[i].Name, name) {
			match = &calendars[i]
		}
	}

	if match == nil {
		// Try partial match
		for i := range calendars {
			if strings.Contains(strings.ToLower(calendars[i].Name), strings.ToLower(name)) {
				match = &calendars[i]
				break
			}
		}
	}

	if match == nil {
		return nil, fmt.Errorf("no calendar found matching '%s'. Available: %s", name, strings.Join(available, ", "))
	}

	return match, nil
}

func disambiguateContact(contacts []models.Contact) (*models.Contact, error) {
	items := make([]string, len(contacts))
	for i, c := range contacts {
		parts := []string{c.DisplayName()}
		if c.Email != "" {
			parts = append(parts, c.Email)
		}
		if c.Phone != "" {
			parts = append(parts, c.Phone)
		}
		items[i] = strings.Join(parts, " | ")
	}

	prompt := promptui.Select{
		Label: "Multiple contacts found. Select one",
		Items: items,
		Size:  10,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("selection cancelled: %w", err)
	}

	return &contacts[idx], nil
}
