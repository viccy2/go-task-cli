package internal

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

// Task represents a single task in our system
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

// TaskList is a slice of Task, we attach methods to this to manage tasks
type TaskList []Task

const fileName = "tasks.json"

// Add creates a new task and appends it to the list
func (t *TaskList) Add(description string) {
	id := 1
	if len(*t) > 0 {
		// Set ID to one higher than the last task's ID
		id = (*t)[len(*t)-1].ID + 1
	}

	task := Task{
		ID:          id,
		Description: description,
		Done:        false,
		CreatedAt:   time.Time{}, // Simplified for now, or use time.Now()
	}
	*t = append(*t, task)
}

// Delete removes a task from the list by its ID
func (t *TaskList) Delete(id int) {
	ls := *t
	for i, task := range ls {
		if task.ID == id {
			*t = append(ls[:i], ls[i+1:]...)
			return
		}
	}
}

// Complete marks a task as done
func (t *TaskList) Complete(id int) {
	ls := *t
	for i := range ls {
		if ls[i].ID == id {
			ls[i].Done = true
			return
		}
	}
}

// Search filters tasks based on a keyword (case-insensitive)
func (t TaskList) Search(keyword string) TaskList {
	results := TaskList{}
	for _, task := range t {
		// strings.Contains checks if the keyword exists inside the description
		if strings.Contains(strings.ToLower(task.Description), strings.ToLower(keyword)) {
			results = append(results, task)
		}
	}
	return results
}

// LoadTasks reads the JSON file and returns a TaskList
func LoadTasks() (TaskList, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return TaskList{}, nil
		}
		return nil, err
	}

	var tasks TaskList
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

// SaveTasks writes the current TaskList to the JSON file
func SaveTasks(tasks TaskList) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
