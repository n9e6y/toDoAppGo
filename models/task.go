package models

import (
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Category    string    `json:"category"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
