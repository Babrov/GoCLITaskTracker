package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const FILE_NAME = "tasks.json"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := loadTasks(FILE_NAME)

	fmt.Println("GO CLI TODO. Type 'help' for commands, or 'exit' to quit.")

	for {
		fmt.Print("task-cli > ")

		if !scanner.Scan() {
			break
		}

		input := strings.Split(scanner.Text(), " ")
		command := input[0]

		if command == "" {
			continue
		}

		switch command {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  list - list all tasks. [todo/in-progress/done]")
			fmt.Println("  add - add a new task with a description. add <description>")
			fmt.Println("  update - update a task. update <id> <description>")
			fmt.Println("  delete - delete a task. delete <id>")
			fmt.Println("  mark-in-progress - mark a task as in-progress. mark-in-progress <id>")
			fmt.Println("  mark-done - mark a task as done. mark-done <id>")
			fmt.Println("  exit - Exit the program")
		case "list":
			args := input[1:]
			if len(args) > 0 && (args[0] == "todo" || args[0] == "in-progress" || args[0] == "done") {
				statusFilter := args[0]
				fmt.Println("Listing tasks with status:", statusFilter)
				for _, task := range tasks.Tasks {
					if task.Status == statusFilter {
						fmt.Println("Id:", task.ID, "Description:", task.Description, "Status:", task.Status)
					}
				}
				continue
			} else {
				fmt.Println("Listing all tasks:")

				for _, task := range tasks.Tasks {
					fmt.Println("Id:", task.ID, "Description:", task.Description, "Status:", task.Status)
				}
			}
		case "add":
			args := input[1:]
			if len(args) > 0 {
				description := args[0]

				task := tasks.add(description)
				fmt.Println("Task", description, "added with ID", task.ID)
			} else {
				fmt.Println("No description.")
			}

			saveTasks(FILE_NAME, tasks)

		case "update":
			args := input[1:]
			if len(args) >= 2 {
				id := args[0]
				description := args[1]

				_, err := tasks.update(id, description)
				if err != nil {
					fmt.Println(err)
				}
				saveTasks(FILE_NAME, tasks)
				fmt.Println("Task", id, "updated.")
			}
		case "delete":
			args := input[1:]
			if len(args) > 0 {
				id := args[0]

				_, err := tasks.delete(id)
				if err != nil {
					fmt.Println(err)
				}
				saveTasks(FILE_NAME, tasks)
				fmt.Println("Task", id, "deleted.")
			}
		case "mark-in-progress":
			args := input[1:]
			if len(args) > 0 {
				id := args[0]

				_, err := tasks.markAs(id, "in-progress")
				if err != nil {
					fmt.Println(err)
				}
				saveTasks(FILE_NAME, tasks)
				fmt.Println("Task", id, "marked as in-progress.")

			}
		case "mark-done":
			args := input[1:]
			if len(args) > 0 {
				id := args[0]

				_, err := tasks.markAs(id, "done")
				if err != nil {
					fmt.Println(err)
				}
				saveTasks(FILE_NAME, tasks)
				fmt.Println("Task", id, "marked as done.")

			}
		case "exit":
			fmt.Println("Exiting CLI. Goodbye!")
			return // Use return to exit the main function
		default:
			fmt.Printf("Unknown command: '%s'. Type 'help' for a list of commands.\n", command)
		}
	}

	// Handle any errors that may have occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from input: %v\n", err)
	}
}

func loadTasks(filename string) TaskList {
	var tasks TaskList

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		emptyTasks := TaskList{Tasks: []TaskItem{}}
		data, _ := json.MarshalIndent(emptyTasks, "", "  ")
		err := os.WriteFile(filename, data, 0644)
		if err != nil {
			return TaskList{}
		}

		return emptyTasks
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &tasks)

	if err != nil {
		return TaskList{}
	}

	return tasks
}

func saveTasks(filename string, tasks TaskList) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return
	}
}
