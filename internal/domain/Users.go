package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `json:"firstname" gorm:"not null"`
	LastName  string `json:"lastname" gorm:"not null"`
	Email     string `json:"email" gorm:"not null; unique_index"`
	Task      []Task `json:"task" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
