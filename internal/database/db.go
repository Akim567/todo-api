package database

import (
	//"fmt"
	"todo-pet/internal/app/task"
	"todo-pet/pkg/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=todo_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// Автоматическая миграция таблицы
	if err := db.AutoMigrate(&task.Todo{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	//fmt.Println("Успешное подключение к базе данных!")
	return db
}
