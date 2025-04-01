package cli

import (
	"fmt"
	"os"
	"strings"

	"todo-pet/internal/app/task"
)

type CLIHandler struct {
	service *task.Service
}

func NewCLIHandler(service *task.Service) *CLIHandler {
	return &CLIHandler{service: service}
}

func (h *CLIHandler) HandleCommand() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Нет команды. Используйте: list")
		return
	}

	switch args[1] {
	case "list":
		h.listTodos()
	default:
		fmt.Println("Неизвестная команда:", args[1])
	}
}

func (h *CLIHandler) listTodos() {
	todos := h.service.GetAll()
	if len(todos) == 0 {
		fmt.Println("Нет заметок")
		return
	}

	fmt.Printf("| %-4s | %-20s | %-10s | %-20s |\n", "ID", "Title", "Status", "Completed At")
	fmt.Println(strings.Repeat("-", 65))

	for _, todo := range todos {
		var status string
		switch todo.Status {
		case "completed":
			status = "completed"
		case "cancelled":
			status = "cancelled"
		default:
			status = "active"
		}

		completedAt := "-"
		if todo.CompletedAt != nil && !todo.CompletedAt.IsZero() {
			completedAt = todo.CompletedAt.Format("02.01.2006 15:04")
		}

		fmt.Printf("| %-4d | %-20s | %-10s | %-20s |\n", todo.ID, todo.Title, status, completedAt)
	}
}
