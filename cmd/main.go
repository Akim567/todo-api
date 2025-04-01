package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"todo-pet/internal/api"
	"todo-pet/internal/app/task"
	"todo-pet/internal/cli"
)

func main() {

	now := time.Now()

	var todos = []task.Todo{
		{ID: 1, Title: "Купить молоко", Status: "active"},
		{ID: 2, Title: "Позвонить другу", Status: "completed", CompletedAt: &now},
		{ID: 3, Title: "Выгулять собаку", Status: "cancelled"},
	}

	service := task.NewService(todos)

	// Если передана CLI-команда — выполняем CLI и выходим
	if len(os.Args) > 1 {
		handlerCLI := cli.NewCLIHandler(service)
		handlerCLI.HandleCommand()
		return
	}

	// Иначе — запускаем API сервер
	handlerAPI := api.NewHandler(service)
	handlerAPI.RegisterRoutes()

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
