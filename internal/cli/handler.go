package cli

import (
	"fmt"
	"os"
	"strconv"
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

	switch args[1] {
	case "list":
		listCmd := &command.ListCommand{Service: h.service}
		if err := listCmd.Execute(); err != nil {
			fmt.Println("Ошибка при выводе списка:", err)
		}

	case "add":
		if len(args) < 3 {
			fmt.Println("Введите название задачи: go run ./cmd add \"Task\"")
			return
		}
		title := strings.Join(args[2:], " ")

		// Сначала добавляем задачу
		addCmd := &command.AddCommand{Service: h.service, Title: title}
		if err := addCmd.Execute(); err != nil {
			fmt.Println("Ошибка при добавлении задачи:", err)
			return
		}

		// Затем выводим список
		listCmd := &command.ListCommand{Service: h.service}
		if err := listCmd.Execute(); err != nil {
			fmt.Println("Ошибка при выводе списка:", err)
		}

	case "delete":
		if len(args) < 3 {
			fmt.Println("Укажите ID задачи для удаления: go run ./cmd delete 2")
			return
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}
		delCmd := &command.DeleteCommand{Service: h.service, ID: id}
		if err := delCmd.Execute(); err != nil {
			fmt.Println("Ошибка при удалении:", err)
			return
		}
		// Можно вывести обновлённый список
		listCmd := &command.ListCommand{Service: h.service}
		_ = listCmd.Execute()

	case "done":
		if len(args) < 3 {
			fmt.Println("Укажите ID задачи для изменения статуса: go run ./cmd delete 2")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}

		doneCmd := &command.DoneCommand{Service: h.service, ID: id}
		if err := doneCmd.Execute(); err != nil {
			fmt.Println("Ошибка при изменении статуса:", err)
			return
		}

		// Можно вывести обновлённый список
		listCmd := &command.ListCommand{Service: h.service}
		_ = listCmd.Execute()

	case "help":
		helpCmd := &command.HelpCommand{}
		_ = helpCmd.Execute()
	default:
		fmt.Println("Неизвестная команды:", args[1])
		return
	}
}
