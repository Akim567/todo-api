package task

import "time"

type Todo struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null"`
	Status      string     `json:"status" gorm:"default:active"`
	CompletedAt *time.Time `json:"completed_at" gorm:"column:completed_at"`
}
