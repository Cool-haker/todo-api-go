package models

import "time"

type Todo struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"user_id" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type TmpTodo struct {
	Title       *string `json:"title" gorm:"not null"`
	Description *string `json:"description"`
	IsCompleted *bool   `json:"is_completed" gorm:"default:false"`
}
