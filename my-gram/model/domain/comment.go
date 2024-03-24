package domain

import (
	"gorm.io/gorm"
)

type Comment struct {
	ID      int `gorm:"primaryKey" json:"id"`
	UserID  int `json:"user_id" gorm:"foreignKey:ID"`
	User    User
	PhotoID int `json:"photo_id" gorm:"foreignKey:ID"`
	Photo   Photo
	Message string `json:"message"`
	gorm.Model
}
