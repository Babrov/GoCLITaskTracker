package main

import (
	"errors"
	"strconv"
	"time"
)

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

func (list *TaskList) getNextID() int {
	maxID := 0
	for _, task := range list.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func (list *TaskList) add(description string) *TaskItem {
	newTask := TaskItem{
		ID:          list.getNextID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: description,
		Status:      "todo",
	}

	list.Tasks = append(list.Tasks, newTask)

	return &newTask
}

func (list *TaskList) get(idStr string) (*TaskItem, error) {
	for i := range list.Tasks {
		task := &list.Tasks[i]
		if idStr == strconv.Itoa(task.ID) {
			return task, nil
		}
	}

	return nil, errors.New("task not found")
}

func (list *TaskList) update(idStr string, description string) (bool, error) {
	task, err := list.get(idStr)

	if err != nil {
		return false, err
	}

	task.Description = description
	task.UpdatedAt = time.Now()

	return true, nil
}

func (list *TaskList) markAs(idStr string, status string) (bool, error) {
	task, err := list.get(idStr)

	if err != nil {
		return false, err
	}

	if status != "todo" && status != "in-progress" && status != "done" {
		return false, errors.New("invalid status")
	}

	task.Status = status
	task.UpdatedAt = time.Now()

	return true, nil
}

func (list *TaskList) delete(idStr string) (bool, error) {
	for i, task := range list.Tasks {
		if idStr == strconv.Itoa(task.ID) {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			return true, nil
		}
	}

	return false, errors.New("task not found")
}
