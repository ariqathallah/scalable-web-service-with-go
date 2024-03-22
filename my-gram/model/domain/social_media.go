package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         string `json:"user_id" gorm:"foreignKey:ID"`
	User           User
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      gorm.DeletedAt
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
