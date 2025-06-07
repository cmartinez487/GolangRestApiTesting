package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `json:"title" gorm:"not nulll; unique_index type:varchar(100)"`
	Description string `json:"description"`
	Done        bool   `json:"done" gorm:"default:false"`
	UserID      uint   `json:"user_id"`
}
