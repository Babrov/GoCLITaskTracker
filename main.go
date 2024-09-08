package main

import (
    "errors"
    "fmt"
    "os"
    "time"
)

type TaskItem struct {
    ID          int       `json:"id"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
}

func CreateTaskItem(id int, description string) (TaskItem, error) {
    if description == "" {
        return TaskItem{}, errors.New("can't create task item with empty description")
    }

    return TaskItem{
        ID:          id,
        Status:      "todo",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
        Description: description,
    }, nil
}

func main() {
    args := os.Args[1:]
    if !(len(args) > 0) {
        fmt.Println("No command provided")

        os.Exit(0)
    }

    command := args[0]

    if command == "add" {
        if len(args) > 1 {
            description := args[1]

            newItem := add(description)

            fmt.Printf("New item created, %v\n", newItem.ID)
        } else {
            fmt.Println("No description.")
        }
    }

}

func add(description string) TaskItem {
    item, err := CreateTaskItem(1, description)

    if err != nil {
        fmt.Println("Error:", err)
    }

    return item
}
