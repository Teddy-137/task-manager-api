package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" binding:"required, max=255"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"required, oneof=pending in_progress completed, default=pending"`
}
