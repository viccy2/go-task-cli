package internal

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskList []Task

// AddTask creates a new task and adds it to the list
func (l *TaskList) Add(description string) {
	newTask := Task{
		ID:          len(*l) + 1,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}
	*l = append(*l, newTask)
}

// Delete removes a task by ID
func (l *TaskList) Delete(id int) {
	updated := TaskList{}
	for _, t := range *l {
		if t.ID != id {
			updated = append(updated, t)
		}
	}
	*l = updated
}
