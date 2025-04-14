package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"todo-pet/internal/api"
	"todo-pet/internal/app/task"
	"todo-pet/internal/cli"
	"todo-pet/internal/database"

	"gorm.io/gorm"
)

func main() {

	db := database.Connect()

	seedDatabaseIfEmpty(db)

	service := task.NewService(db)

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

func seedDatabaseIfEmpty(db *gorm.DB) {
	var count int64
	db.Model(&task.Todo{}).Count(&count)
	if count > 0 {
		return // Уже есть задачи — выходим
	}

	now := time.Now()
	initialTodos := []task.Todo{
		{Title: "Купить молоко", Status: "active"},
		{Title: "Позвонить другу", Status: "completed", CompletedAt: &now},
		{Title: "Выгулять собаку", Status: "cancelled"},
	}

	result := db.Create(&initialTodos)
	if result.Error != nil {
		fmt.Println("Ошибка при инициализации базы:", result.Error)
	} else {
		fmt.Println("Задачи по умолчанию добавлены в базу данных.")
	}
}
