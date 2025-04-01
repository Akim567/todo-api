package cli

import (
	"fmt"
	"os"

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

	fmt.Println("Список задач:")
	for i, todo := range todos {
		status := "[ ]"
		if todo.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, todo.Title)
	}
}
