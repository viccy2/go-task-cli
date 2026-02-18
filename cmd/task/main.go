package main

import (
	"flag"
	"fmt"
	"github.com/viccy2/go-task-cli/internal"
)

func main() {
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	del := flag.Int("del", 0, "Delete a task by ID")

	flag.Parse()

	// We ignore the error for now with '_' to keep it simple
	tasks, _ := internal.LoadTasks()

	if *add != "" {
		tasks.Add(*add)
		internal.SaveTasks(tasks)
		fmt.Println("✅ Task added!")
	} else if *list {
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "X"
			}
			fmt.Printf("[%d] [%s] %s\n", t.ID, status, t.Description)
		}
	} else if *del != 0 {
		tasks.Delete(*del)
		internal.SaveTasks(tasks)
		fmt.Println("🗑️ Task deleted!")
	} else {
		flag.Usage()
	}
}
