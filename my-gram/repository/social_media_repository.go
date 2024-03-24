package repository

import (
	"my-gram/model/domain"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia domain.SocialMedia) (domain.SocialMedia, error)

	GetByID(socialMediaID int) (domain.SocialMedia, error)
	GetAll() ([]domain.SocialMedia, error)

	Update(socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	Delete(socialMediaID int) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) Create(newSocialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Create(&newSocialMedia).Error
	return newSocialMedia, err
}

func (r *socialMediaRepository) GetByID(socialMediaID int) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia
	err := r.db.Where("id = ?", socialMediaID).First(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) GetAll() ([]domain.SocialMedia, error) {
	var socialMedias []domain.SocialMedia
	err := r.db.Find(&socialMedias).Error
	return socialMedias, err
}

func (r *socialMediaRepository) Update(updatedSocialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Save(&updatedSocialMedia).Error
	return updatedSocialMedia, err
}

func (r *socialMediaRepository) Delete(socialMediaID int) error {
	return r.db.Where("id = ?", socialMediaID).Delete(&domain.SocialMedia{}).Error
}
