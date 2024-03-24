package repository

import (
	"my-gram/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user domain.User) (domain.User, error)

	GetByID(userID int) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetByUsername(username string) (domain.User, error)

	Update(user domain.User) (domain.User, error)
	Delete(userID int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(newUser domain.User) (domain.User, error) {
	err := r.db.Create(&newUser).Error
	return newUser, err
}

func (r *userRepository) GetByID(userID int) (domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", userID).First(&user).Error
	return user, err
}

func (r *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) GetByUsername(username string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepository) Update(updatedUser domain.User) (domain.User, error) {
	err := r.db.Save(&updatedUser).Error
	return updatedUser, err
}

func (r *userRepository) Delete(userID int) error {
	return r.db.Where("id = ?", userID).Delete(&domain.User{}).Error
}
