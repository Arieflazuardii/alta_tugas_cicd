package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	UserID  uint   `json:"user_id" form:"user_id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type BlogResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}