package task

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() []Todo {
	var todos []Todo
	s.DB.Find(&todos)
	return todos
}

func (s *Service) Add(title string) Todo {
	todo := Todo{
		Title:  title,
		Status: "active",
	}
	s.DB.Create(&todo)
	return todo
}

func (s *Service) DeleteById(id int) bool {
	result := s.DB.Delete(&Todo{}, id)
	return result.RowsAffected > 0
}

func (s *Service) Done(id int) bool {
	var todo Todo
	result := s.DB.First(&todo, id)
	if result.Error != nil {
		return false
	}

	todo.Status = "completed"
	now := time.Now()
	todo.CompletedAt = &now

	s.DB.Save(&todo)
	return true
}
