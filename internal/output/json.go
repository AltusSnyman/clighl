package output

import (
	"encoding/json"

	"github.com/altusmusic/clighl/internal/models"
)

// JSONFormatter outputs data as formatted JSON.
type JSONFormatter struct{}

func (f *JSONFormatter) FormatContacts(contacts []models.Contact) string {
	return marshalIndent(contacts)
}

func (f *JSONFormatter) FormatContact(contact *models.Contact) string {
	return marshalIndent(contact)
}

func (f *JSONFormatter) FormatPipelines(pipelines []models.Pipeline) string {
	return marshalIndent(pipelines)
}

func (f *JSONFormatter) FormatPipeline(pipeline *models.Pipeline) string {
	return marshalIndent(pipeline)
}

func (f *JSONFormatter) FormatOpportunities(opportunities []models.Opportunity) string {
	return marshalIndent(opportunities)
}

func (f *JSONFormatter) FormatOpportunity(opportunity *models.Opportunity) string {
	return marshalIndent(opportunity)
}

func (f *JSONFormatter) FormatCalendars(calendars []models.Calendar) string {
	return marshalIndent(calendars)
}

func (f *JSONFormatter) FormatCalendar(calendar *models.Calendar) string {
	return marshalIndent(calendar)
}

func (f *JSONFormatter) FormatAppointment(appointment *models.Appointment) string {
	return marshalIndent(appointment)
}

func (f *JSONFormatter) FormatFreeSlots(slots map[string][]string) string {
	return marshalIndent(slots)
}

func (f *JSONFormatter) FormatNotes(notes []models.Note) string {
	return marshalIndent(notes)
}

func (f *JSONFormatter) FormatNote(note *models.Note) string {
	return marshalIndent(note)
}

func (f *JSONFormatter) FormatTags(tags []models.Tag) string {
	return marshalIndent(tags)
}

func (f *JSONFormatter) FormatContactTags(tags []string) string {
	return marshalIndent(tags)
}

func (f *JSONFormatter) FormatTasks(tasks []models.Task) string {
	return marshalIndent(tasks)
}

func (f *JSONFormatter) FormatConversations(convos []models.Conversation) string {
	return marshalIndent(convos)
}

func (f *JSONFormatter) FormatMessages(messages []models.Message) string {
	return marshalIndent(messages)
}

func (f *JSONFormatter) FormatMessage(message *models.Message) string {
	return marshalIndent(message)
}

func (f *JSONFormatter) FormatLocation(location *models.Location) string {
	return marshalIndent(location)
}

func (f *JSONFormatter) FormatCustomFields(fields []models.LocationCustomField) string {
	return marshalIndent(fields)
}

func (f *JSONFormatter) FormatAppointments(appointments []models.Appointment) string {
	return marshalIndent(appointments)
}

func (f *JSONFormatter) FormatBlogs(blogs []models.Blog) string {
	return marshalIndent(blogs)
}
func (f *JSONFormatter) FormatBlogPosts(posts []models.BlogPost) string {
	return marshalIndent(posts)
}
func (f *JSONFormatter) FormatBlogPost(post *models.BlogPost) string {
	return marshalIndent(post)
}
func (f *JSONFormatter) FormatBlogAuthors(authors []models.BlogAuthor) string {
	return marshalIndent(authors)
}
func (f *JSONFormatter) FormatBlogCategories(categories []models.BlogCategory) string {
	return marshalIndent(categories)
}
func (f *JSONFormatter) FormatSocialAccounts(accounts []models.SocialAccount) string {
	return marshalIndent(accounts)
}
func (f *JSONFormatter) FormatSocialPosts(posts []models.SocialPost) string {
	return marshalIndent(posts)
}
func (f *JSONFormatter) FormatSocialPost(post *models.SocialPost) string {
	return marshalIndent(post)
}
func (f *JSONFormatter) FormatEmailTemplates(templates []models.EmailTemplate) string {
	return marshalIndent(templates)
}
func (f *JSONFormatter) FormatEmailTemplate(template *models.EmailTemplate) string {
	return marshalIndent(template)
}
func (f *JSONFormatter) FormatOrder(order *models.Order) string {
	return marshalIndent(order)
}
func (f *JSONFormatter) FormatTransactions(transactions []models.Transaction) string {
	return marshalIndent(transactions)
}

func marshalIndent(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return `{"error": "failed to marshal JSON"}`
	}
	return string(data)
}
