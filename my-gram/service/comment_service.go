package service

import (
	"my-gram/exception"
	"my-gram/helper"
	"my-gram/model/domain"
	"my-gram/model/web"
	"my-gram/repository"
)

type CommentService interface {
	CreateComment(userID int, request web.CreateCommentRequest) (web.CreateCommentResponse, *exception.CustomError)
	GetAllComments() ([]web.GetCommentResponse, *exception.CustomError)
	UpdateComment(userID, commentID int, request web.UpdateCommentRequest) (web.UpdateCommentResponse, *exception.CustomError)
	DeleteComment(userID, commentID int) *exception.CustomError
}

type commentService struct {
	CommentRepo repository.CommentRepository
	PhotoRepo   repository.PhotoRepository
	UserRepo    repository.UserRepository
}

func NewCommentService(commentRepo repository.CommentRepository, photoRepo repository.PhotoRepository, userRepo repository.UserRepository) *commentService {
	return &commentService{
		CommentRepo: commentRepo,
		PhotoRepo:   photoRepo,
		UserRepo:    userRepo,
	}
}

func (s *commentService) CreateComment(userID int, request web.CreateCommentRequest) (web.CreateCommentResponse, *exception.CustomError) {
	// validate request
	validateMessages := helper.CreateCommentValidate(request)
	for _, message := range validateMessages {
		return web.CreateCommentResponse{}, exception.ErrBadRequest(message)
	}

	// check if photo exists
	photo, err := s.PhotoRepo.GetByID(request.PhotoID)
	if err != nil {
		return web.CreateCommentResponse{}, exception.ErrNotFound("Photo not found")
	}

	// create comment
	comment := domain.Comment{
		UserID:  userID,
		PhotoID: photo.ID,
		Message: request.Message,
	}

	// insert comment to database
	savedComment, err := s.CommentRepo.Create(comment)
	if err != nil {
		return web.CreateCommentResponse{}, exception.ErrInternalServer("Failed to save comment to database")
	}

	response := web.CreateCommentResponse{
		ID:        savedComment.ID,
		Message:   savedComment.Message,
		PhotoID:   savedComment.PhotoID,
		UserID:    savedComment.UserID,
		CreatedAt: savedComment.CreatedAt,
	}

	return response, nil
}

func (s *commentService) GetAllComments() ([]web.GetCommentResponse, *exception.CustomError) {
	// get all comments from database
	comments, err := s.CommentRepo.GetAll()
	if err != nil {
		return nil, exception.ErrInternalServer("Failed to get comments from database")
	}

	var response []web.GetCommentResponse
	for _, comment := range comments {
		// retrieve user data
		user, err := s.UserRepo.GetByID(comment.UserID)
		if err != nil {
			return nil, exception.ErrInternalServer("Failed to get user data from database")
		}

		// retrieve photo data
		photo, err := s.PhotoRepo.GetByID(comment.PhotoID)
		if err != nil {
			return nil, exception.ErrInternalServer("Failed to get photo data from database")
		}

		response = append(response, web.GetCommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User: web.UserGetCommentResponse{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			},
			Photo: web.PhotoGetCommentResponse{
				ID:       photo.ID,
				Title:    photo.Title,
				Caption:  photo.Caption,
				PhotoURL: photo.PhotoURL,
				UserID:   photo.UserID,
			},
		})
	}

	return response, nil
}

func (s *commentService) UpdateComment(userID, commentID int, request web.UpdateCommentRequest) (web.UpdateCommentResponse, *exception.CustomError) {
	// validate request
	validateMessages := helper.UpdateCommentValidate(request)
	for _, message := range validateMessages {
		return web.UpdateCommentResponse{}, exception.ErrBadRequest(message)
	}

	// check if comment exists
	comment, err := s.CommentRepo.GetByID(commentID)
	if err != nil {
		return web.UpdateCommentResponse{}, exception.ErrNotFound("Comment not found")
	}

	// check if the comment belongs to the user
	if comment.UserID != userID {
		return web.UpdateCommentResponse{}, exception.ErrForbidden("You are not allowed to update this comment")
	}

	// update comment
	comment.Message = request.Message

	// save comment to database
	updatedComment, err := s.CommentRepo.Update(comment)
	if err != nil {
		return web.UpdateCommentResponse{}, exception.ErrInternalServer("Failed to update comment to database")
	}

	// get photo data
	photo, err := s.PhotoRepo.GetByID(updatedComment.PhotoID)
	if err != nil {
		return web.UpdateCommentResponse{}, exception.ErrInternalServer("Failed to get photo data from database")
	}

	response := web.UpdateCommentResponse{
		ID:        updatedComment.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: updatedComment.UpdatedAt,
	}

	return response, nil
}

func (s *commentService) DeleteComment(userID, commentID int) *exception.CustomError {
	// check if comment exists
	comment, err := s.CommentRepo.GetByID(commentID)
	if err != nil {
		return exception.ErrNotFound("Comment not found")
	}

	// check if the comment belongs to the user
	if comment.UserID != userID {
		return exception.ErrForbidden("You are not allowed to delete this comment")
	}

	// delete comment from database
	if err := s.CommentRepo.Delete(comment.ID); err != nil {
		return exception.ErrInternalServer("Failed to delete comment from database")
	}

	return nil
}
