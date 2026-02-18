package main

import (
	"flag"
	"fmt"
	"github.com/viccy2/go-task-cli/internal"
)

func main() {
	add := flag.String("add", "", "Add a task")
	list := flag.Bool("list", false, "List all tasks")
	del := flag.Int("del", 0, "Delete task by ID")
	done := flag.Int("done", 0, "Mark task as complete by ID")

	flag.Parse()

	tasks, _ := internal.LoadTasks()

	switch {
	case *add != "":
		tasks.Add(*add)
		internal.SaveTasks(tasks)
		fmt.Println("✅ Task added!")

	case *list:
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Printf("%-3s %-20s %-10s %-15s\n", "ID", "Task", "Status", "Created")
		for _, t := range tasks {
			status := "Pending"
			if t.Done {
				status = "Done"
			}
			// Format the time to a readable string
			created := t.CreatedAt.Format("Jan 02 15:04")
			fmt.Printf("%-3d %-20s %-10s %-15s\n", t.ID, t.Description, status, created)
		}

	case *del != 0:
		tasks.Delete(*del)
		internal.SaveTasks(tasks)
		fmt.Println("🗑️ Task deleted!")

	case *done != 0:
		tasks.Complete(*done)
		internal.SaveTasks(tasks)
		fmt.Println("✔️ Task marked as complete!")

	default:
		flag.Usage()
	}
}
