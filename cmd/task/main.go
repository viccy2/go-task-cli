package main

import (
	"flag"
	"fmt"
	"github.com/viccy2/go-task-cli/internal"
)

func main() {
	// 1. Define the Flags
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	del := flag.Int("del", 0, "Delete task by ID")
	done := flag.Int("done", 0, "Mark task as complete by ID")
	find := flag.String("find", "", "Search tasks by keyword")

	flag.Parse()

	// 2. Load existing tasks from the JSON file
	tasks, err := internal.LoadTasks()
	if err != nil {
		// If the file doesn't exist yet, we start with an empty list
		tasks = internal.TaskList{}
	}

	// 3. Handle the User's Input
	switch {
	case *add != "":
		tasks.Add(*add)
		internal.SaveTasks(tasks)
		fmt.Println("✅ Task added successfully!")

	case *list:
		if len(tasks) == 0 {
			fmt.Println("📝 Your task list is empty.")
			return
		}
		printHeader()
		for _, t := range tasks {
			printTask(t)
		}

	case *find != "":
		results := tasks.Search(*find)
		if len(results) == 0 {
			fmt.Printf("🔍 No tasks found matching: '%s'\n", *find)
			return
		}
		fmt.Printf("🔍 Search results for '%s':\n", *find)
		printHeader()
		for _, t := range results {
			printTask(t)
		}

	case *del != 0:
		tasks.Delete(*del)
		internal.SaveTasks(tasks)
		fmt.Printf("🗑️ Task %d deleted!\n", *del)

	case *done != 0:
		tasks.Complete(*done)
		internal.SaveTasks(tasks)
		fmt.Printf("✔️ Task %d marked as done!\n", *done)

	default:
		// If no flags are provided, show the help menu
		flag.Usage()
	}
}

// Helper function to keep the UI consistent
func printHeader() {
	fmt.Printf("%-3s %-20s %-10s %-15s\n", "ID", "Task", "Status", "Created")
	fmt.Println("------------------------------------------------------------")
}

// Helper function to format the output of a single task
func printTask(t internal.Task) {
	status := "Pending"
	if t.Done {
		status = "Done"
	}
	created := t.CreatedAt.Format("Jan 02 15:04")
	fmt.Printf("%-3d %-20s %-10s %-15s\n", t.ID, t.Description, status, created)
}
