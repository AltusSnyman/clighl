package output

import (
	"github.com/altusmusic/clighl/internal/models"
)

// Formatter defines how to render CLI output.
type Formatter interface {
	FormatContacts(contacts []models.Contact) string
	FormatContact(contact *models.Contact) string
	FormatPipelines(pipelines []models.Pipeline) string
	FormatPipeline(pipeline *models.Pipeline) string
	FormatOpportunities(opportunities []models.Opportunity) string
	FormatOpportunity(opportunity *models.Opportunity) string
	FormatCalendars(calendars []models.Calendar) string
	FormatCalendar(calendar *models.Calendar) string
	FormatAppointment(appointment *models.Appointment) string
	FormatFreeSlots(slots map[string][]string) string
	FormatNotes(notes []models.Note) string
	FormatNote(note *models.Note) string
	FormatTags(tags []models.Tag) string
	FormatContactTags(tags []string) string
	FormatTasks(tasks []models.Task) string
	FormatConversations(convos []models.Conversation) string
	FormatMessages(messages []models.Message) string
	FormatMessage(message *models.Message) string
	FormatLocation(location *models.Location) string
	FormatCustomFields(fields []models.LocationCustomField) string
	FormatAppointments(appointments []models.Appointment) string
	FormatBlogs(blogs []models.Blog) string
	FormatBlogPosts(posts []models.BlogPost) string
	FormatBlogPost(post *models.BlogPost) string
	FormatBlogAuthors(authors []models.BlogAuthor) string
	FormatBlogCategories(categories []models.BlogCategory) string
	FormatSocialAccounts(accounts []models.SocialAccount) string
	FormatSocialPosts(posts []models.SocialPost) string
	FormatSocialPost(post *models.SocialPost) string
	FormatEmailTemplates(templates []models.EmailTemplate) string
	FormatEmailTemplate(template *models.EmailTemplate) string
	FormatOrder(order *models.Order) string
	FormatTransactions(transactions []models.Transaction) string
}
