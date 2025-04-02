package command

import (
	"fmt"
	"strings"
	"todo-pet/internal/app/task"
)

type ListCommand struct {
	Service *task.Service
}

func (c *ListCommand) Execute() error {
	todos := c.Service.GetAll()
	if len(todos) == 0 {
		fmt.Println("Нет заметок")
		return nil
	}

	fmt.Printf("| %-4s | %-20s | %-10s | %-20s |\n", "ID", "Title", "Status", "Completed At")
	fmt.Println(strings.Repeat("-", 65))

	for _, todo := range todos {
		status := todo.Status
		completedAt := "-"
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format("02.01.2006 15:04")
		}
		fmt.Printf("| %-4d | %-20s | %-10s | %-20s |\n", todo.ID, todo.Title, status, completedAt)
	}

	return nil
}
