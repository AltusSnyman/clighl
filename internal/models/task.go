package models

// Task represents a GHL contact task.
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Completed   bool   `json:"completed"`
	ContactID   string `json:"contactId"`
	AssignedTo  string `json:"assignedTo"`
	DueDate     string `json:"dueDate"`
	DateAdded   string `json:"dateAdded"`
	DateUpdated string `json:"dateUpdated"`
}

// TasksResponse is the response from GET /contacts/{id}/tasks.
type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}
