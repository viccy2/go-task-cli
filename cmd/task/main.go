package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color" // NEW: Import the color package
	"github.com/viccy2/go-task-cli/internal"
)

func main() {
	add := flag.String("add", "", "Add a task")
	list := flag.Bool("list", false, "List all tasks")
	del := flag.Int("del", 0, "Delete task by ID")
	done := flag.Int("done", 0, "Mark task as complete by ID")

	flag.Parse()

	tasks, _ := internal.LoadTasks()

	// Create color functions
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgCyan).BoldFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	switch {
	case *add != "":
		tasks.Add(*add)
		internal.SaveTasks(tasks)
		fmt.Printf("%s Task added!\n", green("✅"))

	case *list:
		if len(tasks) == 0 {
			fmt.Println(yellow("No tasks found. Get to work!"))
			return
		}
		
		fmt.Printf("%-3s %-20s %-10s %-15s\n", blue("ID"), blue("Task"), blue("Status"), blue("Created"))
		
		for _, t := range tasks {
			status := yellow("Pending")
			if t.Done {
				status = green("Done")
			}
			created := t.CreatedAt.Format("Jan 02 15:04")
			fmt.Printf("%-3d %-20s %-10s %-15s\n", t.ID, t.Description, status, magenta(created))
		}

	case *del != 0:
		tasks.Delete(*del)
		internal.SaveTasks(tasks)
		fmt.Printf("%s Task deleted!\n", yellow("🗑️"))

	case *done != 0:
		tasks.Complete(*done)
		internal.SaveTasks(tasks)
		fmt.Printf("%s Task marked as complete!\n", green("✔️"))

	default:
		flag.Usage()
	}
}
