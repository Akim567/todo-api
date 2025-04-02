package cli

import (
	"fmt"
	"os"
	"strings"

	"todo-pet/internal/app/command"
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

	var cmd command.Command

	switch args[1] {
	case "list":
		cmd = &command.ListCommand{Service: h.service}
	case "add":
		if len(args) < 3 {
			fmt.Println("Введите название задачи: go run ./cmd add \"Task\"")
			return
		}
		title := strings.Join(args[2:], " ")
		cmd = &command.AddCommand{Service: h.service, Title: title}
	default:
		fmt.Println("Неизвестная команды:", args[1])
		return
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println("Ошибка при выполнении команды", err)
	}
}
