package repository

import (
	"my-gram/model/domain"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment domain.Comment) (domain.Comment, error)

	GetByID(commentID int) (domain.Comment, error)
	GetAll() ([]domain.Comment, error)

	Update(comment domain.Comment) (domain.Comment, error)
	Delete(commentID int) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Create(newComment domain.Comment) (domain.Comment, error) {
	err := r.db.Create(&newComment).Error
	return newComment, err
}

func (r *commentRepository) GetByID(commentID int) (domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Where("id = ?", commentID).First(&comment).Error
	return comment, err
}

func (r *commentRepository) GetAll() ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.db.Find(&comments).Error
	return comments, err
}

func (r *commentRepository) Update(updatedComment domain.Comment) (domain.Comment, error) {
	err := r.db.Save(&updatedComment).Error
	return updatedComment, err
}

func (r *commentRepository) Delete(commentID int) error {
	return r.db.Where("id = ?", commentID).Delete(&domain.Comment{}).Error
}
