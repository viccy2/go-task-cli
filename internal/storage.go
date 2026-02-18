package internal

import (
	"encoding/json"
	"os" // Used below
)

const FileName = "tasks.json"

func SaveTasks(tasks TaskList) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	// os used here
	return os.WriteFile(FileName, data, 0644)
}

func LoadTasks() (TaskList, error) {
	// os used here
	data, err := os.ReadFile(FileName)
	if err != nil {
		// os used here
		if os.IsNotExist(err) {
			return TaskList{}, nil
		}
		return nil, err
	}
	var tasks TaskList
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
