package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/viccy2/go-task-cli/internal" 
)

func main() {
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	del := flag.Int("del", 0, "Delete a task by ID")

	flag.Parse()

	tasks, _ := internal.LoadTasks()

	if *add != "" {
		tasks.Add(*add)
		internal.SaveTasks(tasks)
		fmt.Println("Added!")
	} else if *list {
		for _, t := range tasks {
			fmt.Printf("[%d] %s (Done: %v)\n", t.ID, t.Description, t.Done)
		}
	} else if *del != 0 {
		tasks.Delete(*del)
		internal.SaveTasks(tasks)
		fmt.Println("Deleted!")
	} else {
		flag.Usage()
	}
}
