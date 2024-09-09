package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "time"
)

const filename = "tasks.json"

type TaskItem struct {
    ID          int       `json:"id"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
}

type TaskList struct {
    Tasks []TaskItem `json:"tasks"`
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

func loadTasks() TaskList {
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

    errr := json.Unmarshal(file, &tasks)

    if errr != nil {
        return TaskList{}
    }

    return tasks
}

func saveTasks(tasks TaskList) {
    data, _ := json.MarshalIndent(tasks, "", "  ")
    err := os.WriteFile(filename, data, 0644)
    if err != nil {
        return
    }
}

func getNextID(tasks TaskList) int {
    maxID := 0
    for _, task := range tasks.Tasks {
        if task.ID > maxID {
            maxID = task.ID
        }
    }
    return maxID + 1
}

func add(description string) {
    tasks := loadTasks()

    newTask := TaskItem{
        ID:          getNextID(tasks),
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
        Description: description,
        Status:      "todo",
    }

    tasks.Tasks = append(tasks.Tasks, newTask)
    saveTasks(tasks)

    fmt.Printf("Task '%s' added with ID %d\n", description, newTask.ID)
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

            add(description)
        } else {
            fmt.Println("No description.")
        }
    }

}
