package internal

import "time"

// Task represents a single todo item with metadata
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

// TaskList is a collection of Tasks with helper methods
type TaskList []Task

// Add creates a new task and appends it to the list
func (l *TaskList) Add(description string) {
	newTask := Task{
		ID:          len(*l) + 1,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(), // Captures current time
	}
	*l = append(*l, newTask)
}

// Delete removes a task by its ID and returns a new list
func (l *TaskList) Delete(id int) {
	updated := TaskList{}
	for _, t := range *l {
		if t.ID != id {
			updated = append(updated, t)
		}
	}
	*l = updated
}

// Complete finds a task by ID and marks it as done
func (l *TaskList) Complete(id int) {
	for i, t := range *l {
		if t.ID == id {
			(*l)[i].Done = true
			break
		}
	}
}
