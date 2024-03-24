package domain

import (
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             int    `gorm:"primaryKey" json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         int    `json:"user_id" gorm:"foreignKey:ID"`
	User           User
	gorm.Model
}
