package repository

import (
	"my-gram/model/domain"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo domain.Photo) (domain.Photo, error)

	GetByID(photoID int) (domain.Photo, error)
	GetByUserID(userID int) ([]domain.Photo, error)
	GetAll() ([]domain.Photo, error)

	Update(photo domain.Photo) (domain.Photo, error)
	Delete(photoID int) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) Create(newPhoto domain.Photo) (domain.Photo, error) {
	err := r.db.Create(&newPhoto).Error
	return newPhoto, err
}

func (r *photoRepository) GetByID(photoID int) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Where("id = ?", photoID).First(&photo).Error
	return photo, err
}

func (r *photoRepository) GetByUserID(userID int) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.Where("user_id = ?", userID).Find(&photos).Error
	return photos, err
}

func (r *photoRepository) GetAll() ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.Find(&photos).Error
	return photos, err
}

func (r *photoRepository) Update(updatedPhoto domain.Photo) (domain.Photo, error) {
	err := r.db.Save(&updatedPhoto).Error
	return updatedPhoto, err
}

func (r *photoRepository) Delete(photoID int) error {
	return r.db.Where("id = ?", photoID).Delete(&domain.Photo{}).Error
}
