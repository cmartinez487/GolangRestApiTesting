package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `json:"title" gorm:"not nulll; unique_index"`
	Description string
	Done        bool `json:"done" gorm:"default:false"`
	UserID      uint
}
