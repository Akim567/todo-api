package main

import (
	"fmt"
	"net/http"
	"os"

	"todo-pet/internal/api"
	"todo-pet/internal/app/task"
	"todo-pet/internal/cli"
)

func main() {

	var todos = []task.Todo{
		{Title: "Купить молоко", Completed: false},
		{Title: "Позвонить другу", Completed: true},
		{Title: "Выгулять собаку", Completed: false},
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
