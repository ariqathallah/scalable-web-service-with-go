package domain

import (
	"gorm.io/gorm"
)

type Photo struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id" gorm:"foreignKey:ID"`
	User     User
	gorm.Model
}
