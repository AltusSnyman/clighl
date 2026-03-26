package output

import (
	"fmt"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/altusmusic/clighl/internal/models"
)

// TableFormatter outputs data as human-readable tables.
type TableFormatter struct{}

func (f *TableFormatter) FormatContacts(contacts []models.Contact) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tEMAIL\tPHONE\tCOMPANY")
	fmt.Fprintln(w, "──\t────\t─────\t─────\t───────")
	for _, c := range contacts {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			c.ID, c.DisplayName(), c.Email, c.Phone, c.CompanyName)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatContact(contact *models.Contact) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:         %s\n", contact.ID)
	fmt.Fprintf(&buf, "Name:       %s\n", contact.DisplayName())
	fmt.Fprintf(&buf, "Email:      %s\n", contact.Email)
	fmt.Fprintf(&buf, "Phone:      %s\n", contact.Phone)
	fmt.Fprintf(&buf, "Company:    %s\n", contact.CompanyName)
	fmt.Fprintf(&buf, "Source:     %s\n", contact.Source)
	if len(contact.Tags) > 0 {
		fmt.Fprintf(&buf, "Tags:       %s\n", strings.Join(contact.Tags, ", "))
	}
	fmt.Fprintf(&buf, "Added:      %s\n", contact.DateAdded)
	return buf.String()
}

func (f *TableFormatter) FormatPipelines(pipelines []models.Pipeline) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tSTAGES")
	fmt.Fprintln(w, "──\t────\t──────")
	for _, p := range pipelines {
		stageNames := make([]string, len(p.Stages))
		for i, s := range p.Stages {
			stageNames[i] = s.Name
		}
		fmt.Fprintf(w, "%s\t%s\t%s\n", p.ID, p.Name, strings.Join(stageNames, " → "))
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatPipeline(pipeline *models.Pipeline) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:     %s\n", pipeline.ID)
	fmt.Fprintf(&buf, "Name:   %s\n", pipeline.Name)
	fmt.Fprintf(&buf, "Stages:\n")
	for i, s := range pipeline.Stages {
		fmt.Fprintf(&buf, "  %d. %s (ID: %s)\n", i+1, s.Name, s.ID)
	}
	return buf.String()
}

func (f *TableFormatter) FormatOpportunities(opportunities []models.Opportunity) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tSTATUS\tVALUE\tCONTACT ID")
	fmt.Fprintln(w, "──\t────\t──────\t─────\t──────────")
	for _, o := range opportunities {
		fmt.Fprintf(w, "%s\t%s\t%s\t%.2f\t%s\n",
			o.ID, o.Name, o.Status, o.MonetaryValue, o.ContactID)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatOpportunity(opportunity *models.Opportunity) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:         %s\n", opportunity.ID)
	fmt.Fprintf(&buf, "Name:       %s\n", opportunity.Name)
	fmt.Fprintf(&buf, "Status:     %s\n", opportunity.Status)
	fmt.Fprintf(&buf, "Value:      %.2f\n", opportunity.MonetaryValue)
	fmt.Fprintf(&buf, "Pipeline:   %s\n", opportunity.PipelineID)
	fmt.Fprintf(&buf, "Stage:      %s\n", opportunity.PipelineStageID)
	fmt.Fprintf(&buf, "Contact:    %s\n", opportunity.ContactID)
	return buf.String()
}

func (f *TableFormatter) FormatCalendars(calendars []models.Calendar) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDURATION\tACTIVE")
	fmt.Fprintln(w, "──\t────\t────────\t──────")
	for _, c := range calendars {
		active := "no"
		if c.IsActive {
			active = "yes"
		}
		duration := ""
		if c.SlotDuration > 0 {
			duration = fmt.Sprintf("%d min", c.SlotDuration)
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", c.ID, c.Name, duration, active)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatCalendar(calendar *models.Calendar) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:           %s\n", calendar.ID)
	fmt.Fprintf(&buf, "Name:         %s\n", calendar.Name)
	fmt.Fprintf(&buf, "Description:  %s\n", calendar.Description)
	if calendar.SlotDuration > 0 {
		fmt.Fprintf(&buf, "Duration:     %d min\n", calendar.SlotDuration)
	}
	if calendar.SlotInterval > 0 {
		fmt.Fprintf(&buf, "Interval:     %d min\n", calendar.SlotInterval)
	}
	fmt.Fprintf(&buf, "Active:       %v\n", calendar.IsActive)
	return buf.String()
}

func (f *TableFormatter) FormatAppointment(appointment *models.Appointment) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:         %s\n", appointment.ID)
	fmt.Fprintf(&buf, "Calendar:   %s\n", appointment.CalendarID)
	fmt.Fprintf(&buf, "Contact:    %s\n", appointment.ContactID)
	fmt.Fprintf(&buf, "Title:      %s\n", appointment.Title)
	fmt.Fprintf(&buf, "Status:     %s\n", appointment.AppointmentStatus)
	fmt.Fprintf(&buf, "Slot:       %s\n", appointment.SelectedSlot)
	fmt.Fprintf(&buf, "Timezone:   %s\n", appointment.SelectedTimezone)
	if appointment.Notes != "" {
		fmt.Fprintf(&buf, "Notes:      %s\n", appointment.Notes)
	}
	return buf.String()
}

func (f *TableFormatter) FormatFreeSlots(slots map[string][]string) string {
	var buf strings.Builder

	// Sort dates for consistent output
	dates := make([]string, 0, len(slots))
	for date := range slots {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	for _, date := range dates {
		times := slots[date]
		if len(times) == 0 {
			continue
		}
		fmt.Fprintf(&buf, "%s:\n", date)
		for _, t := range times {
			fmt.Fprintf(&buf, "  %s\n", t)
		}
	}
	return buf.String()
}

func (f *TableFormatter) FormatNotes(notes []models.Note) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tDATE\tBODY")
	fmt.Fprintln(w, "──\t────\t────")
	for _, n := range notes {
		body := n.Body
		if len(body) > 60 {
			body = body[:57] + "..."
		}
		fmt.Fprintf(w, "%s\t%s\t%s\n", n.ID, n.DateAdded, body)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatNote(note *models.Note) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:      %s\n", note.ID)
	fmt.Fprintf(&buf, "Date:    %s\n", note.DateAdded)
	fmt.Fprintf(&buf, "Body:    %s\n", note.Body)
	return buf.String()
}

func (f *TableFormatter) FormatTags(tags []models.Tag) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME")
	fmt.Fprintln(w, "──\t────")
	for _, t := range tags {
		fmt.Fprintf(w, "%s\t%s\n", t.ID, t.Name)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatContactTags(tags []string) string {
	if len(tags) == 0 {
		return "No tags\n"
	}
	return strings.Join(tags, ", ") + "\n"
}

func (f *TableFormatter) FormatTasks(tasks []models.Task) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTITLE\tDUE\tDONE\tASSIGNED TO")
	fmt.Fprintln(w, "──\t─────\t───\t────\t───────────")
	for _, t := range tasks {
		done := "no"
		if t.Completed {
			done = "yes"
		}
		title := t.Title
		if title == "" {
			title = t.Body
		}
		if len(title) > 50 {
			title = title[:47] + "..."
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", t.ID, title, t.DueDate, done, t.AssignedTo)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatConversations(convos []models.Conversation) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tCONTACT\tTYPE\tLAST MESSAGE")
	fmt.Fprintln(w, "──\t───────\t────\t────────────")
	for _, c := range convos {
		name := c.FullName
		if name == "" {
			name = c.ContactName
		}
		if name == "" {
			name = c.Email
		}
		if name == "" {
			name = c.Phone
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", c.ID, name, c.Type, c.LastMessageDate())
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatMessages(messages []models.Message) string {
	var buf strings.Builder
	for _, m := range messages {
		dir := "→"
		if m.Direction == "inbound" {
			dir = "←"
		}
		body := m.Body
		if len(body) > 80 {
			body = body[:77] + "..."
		}
		fmt.Fprintf(&buf, "%s [%s] %s %s\n", m.DateAdded, m.Type, dir, body)
	}
	return buf.String()
}

func (f *TableFormatter) FormatMessage(message *models.Message) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:          %s\n", message.ID)
	fmt.Fprintf(&buf, "Type:        %s\n", message.Type)
	fmt.Fprintf(&buf, "Direction:   %s\n", message.Direction)
	fmt.Fprintf(&buf, "Status:      %s\n", message.Status)
	fmt.Fprintf(&buf, "Body:        %s\n", message.Body)
	fmt.Fprintf(&buf, "Date:        %s\n", message.DateAdded)
	return buf.String()
}

func (f *TableFormatter) FormatLocation(location *models.Location) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:        %s\n", location.ID)
	fmt.Fprintf(&buf, "Name:      %s\n", location.Name)
	fmt.Fprintf(&buf, "Email:     %s\n", location.Email)
	fmt.Fprintf(&buf, "Phone:     %s\n", location.Phone)
	fmt.Fprintf(&buf, "Website:   %s\n", location.Website)
	fmt.Fprintf(&buf, "Timezone:  %s\n", location.Timezone)
	addr := location.Address
	if location.City != "" {
		addr += ", " + location.City
	}
	if location.State != "" {
		addr += ", " + location.State
	}
	if location.PostalCode != "" {
		addr += " " + location.PostalCode
	}
	if location.Country != "" {
		addr += ", " + location.Country
	}
	if addr != "" {
		fmt.Fprintf(&buf, "Address:   %s\n", addr)
	}
	return buf.String()
}

func (f *TableFormatter) FormatCustomFields(fields []models.LocationCustomField) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tKEY\tTYPE")
	fmt.Fprintln(w, "──\t────\t───\t────")
	for _, cf := range fields {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", cf.ID, cf.Name, cf.FieldKey, cf.DataType)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatAppointments(appointments []models.Appointment) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTITLE\tSTATUS\tSLOT\tCONTACT ID")
	fmt.Fprintln(w, "──\t─────\t──────\t────\t──────────")
	for _, a := range appointments {
		slot := a.SelectedSlot
		if slot == "" {
			slot = a.StartTime
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", a.ID, a.Title, a.AppointmentStatus, slot, a.ContactID)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatBlogs(blogs []models.Blog) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tURL")
	fmt.Fprintln(w, "──\t────\t───")
	for _, b := range blogs {
		fmt.Fprintf(w, "%s\t%s\t%s\n", b.ID, b.Name, b.URL)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatBlogPosts(posts []models.BlogPost) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTITLE\tSLUG\tSTATUS\tAUTHOR")
	fmt.Fprintln(w, "──\t─────\t────\t──────\t──────")
	for _, p := range posts {
		title := p.Title
		if len(title) > 40 {
			title = title[:37] + "..."
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", p.ID, title, p.Slug, p.Status, p.Author)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatBlogPost(post *models.BlogPost) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:       %s\n", post.ID)
	fmt.Fprintf(&buf, "Title:    %s\n", post.Title)
	fmt.Fprintf(&buf, "Slug:     %s\n", post.Slug)
	fmt.Fprintf(&buf, "Status:   %s\n", post.Status)
	fmt.Fprintf(&buf, "Author:   %s\n", post.Author)
	fmt.Fprintf(&buf, "Category: %s\n", post.CategoryID)
	if len(post.Tags) > 0 {
		fmt.Fprintf(&buf, "Tags:     %s\n", strings.Join(post.Tags, ", "))
	}
	fmt.Fprintf(&buf, "Added:    %s\n", post.DateAdded)
	return buf.String()
}

func (f *TableFormatter) FormatBlogAuthors(authors []models.BlogAuthor) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME")
	fmt.Fprintln(w, "──\t────")
	for _, a := range authors {
		fmt.Fprintf(w, "%s\t%s\n", a.ID, a.Name)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatBlogCategories(categories []models.BlogCategory) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME")
	fmt.Fprintln(w, "──\t────")
	for _, c := range categories {
		fmt.Fprintf(w, "%s\t%s\n", c.ID, c.Name)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatSocialAccounts(accounts []models.SocialAccount) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tPLATFORM\tTYPE")
	fmt.Fprintln(w, "──\t────\t────────\t────")
	for _, a := range accounts {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", a.ID, a.Name, a.Platform, a.Type)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatSocialPosts(posts []models.SocialPost) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tPLATFORM\tSTATUS\tSCHEDULED\tSUMMARY")
	fmt.Fprintln(w, "──\t────────\t──────\t─────────\t───────")
	for _, p := range posts {
		summary := p.Summary
		if summary == "" {
			summary = p.Content
		}
		if len(summary) > 40 {
			summary = summary[:37] + "..."
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", p.ID, p.Platform, p.Status, p.ScheduledAt, summary)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatSocialPost(post *models.SocialPost) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:          %s\n", post.ID)
	fmt.Fprintf(&buf, "Platform:    %s\n", post.Platform)
	fmt.Fprintf(&buf, "Status:      %s\n", post.Status)
	fmt.Fprintf(&buf, "Scheduled:   %s\n", post.ScheduledAt)
	fmt.Fprintf(&buf, "Published:   %s\n", post.PublishedAt)
	fmt.Fprintf(&buf, "Content:     %s\n", post.Content)
	return buf.String()
}

func (f *TableFormatter) FormatEmailTemplates(templates []models.EmailTemplate) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tSUBJECT")
	fmt.Fprintln(w, "──\t────\t───────")
	for _, t := range templates {
		subject := t.Subject
		if len(subject) > 50 {
			subject = subject[:47] + "..."
		}
		fmt.Fprintf(w, "%s\t%s\t%s\n", t.ID, t.Name, subject)
	}
	w.Flush()
	return buf.String()
}

func (f *TableFormatter) FormatEmailTemplate(template *models.EmailTemplate) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:       %s\n", template.ID)
	fmt.Fprintf(&buf, "Name:     %s\n", template.Name)
	fmt.Fprintf(&buf, "Subject:  %s\n", template.Subject)
	fmt.Fprintf(&buf, "Added:    %s\n", template.DateAdded)
	return buf.String()
}

func (f *TableFormatter) FormatOrder(order *models.Order) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "ID:       %s\n", order.ID)
	fmt.Fprintf(&buf, "Name:     %s\n", order.Name)
	fmt.Fprintf(&buf, "Status:   %s\n", order.Status)
	fmt.Fprintf(&buf, "Amount:   %v %s\n", order.Amount, order.Currency)
	fmt.Fprintf(&buf, "Contact:  %s (%s)\n", order.ContactName, order.ContactID)
	fmt.Fprintf(&buf, "Created:  %v\n", order.DateAdded)
	return buf.String()
}

func (f *TableFormatter) FormatTransactions(transactions []models.Transaction) string {
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tAMOUNT\tCURRENCY\tSTATUS\tTYPE\tCONTACT")
	fmt.Fprintln(w, "──\t──────\t────────\t──────\t────\t───────")
	for _, t := range transactions {
		fmt.Fprintf(w, "%s\t%v\t%s\t%s\t%s\t%s\n", t.ID, t.Amount, t.Currency, t.Status, t.Type, t.ContactName)
	}
	w.Flush()
	return buf.String()
}
