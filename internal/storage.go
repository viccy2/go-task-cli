package internal

import (
	"encoding/json"
	"os"
)

const FileName = "tasks.json"

func SaveTasks(tasks TaskList) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(FileName, data, 0644)
}

func LoadTasks() (TaskList, error) {
	data, err := os.ReadFile(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return TaskList{}, nil
		}
		return nil, err
	}
	var tasks TaskList
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
