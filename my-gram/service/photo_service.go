package service

import (
	"my-gram/exception"
	"my-gram/model/domain"
	"my-gram/model/web"
	"my-gram/repository"

	"github.com/go-playground/validator"
)

type PhotoService interface {
	CreatePhoto(userID int, request web.PhotoRequest) (web.CreatePhotoResponse, *exception.CustomError)
	GetAllPhotos() ([]web.GetPhotoResponse, *exception.CustomError)
	UpdatePhoto(userID, photoID int, request web.PhotoRequest) (web.UpdatePhotoResponse, *exception.CustomError)
	DeletePhoto(userID, photoID int) *exception.CustomError
}

type photoService struct {
	Validate  *validator.Validate
	PhotoRepo repository.PhotoRepository
	UserRepo  repository.UserRepository
}

func NewPhotoService(validate *validator.Validate, photoRepo repository.PhotoRepository, userRepo repository.UserRepository) *photoService {
	return &photoService{
		Validate:  validate,
		PhotoRepo: photoRepo,
		UserRepo:  userRepo,
	}
}

func (s *photoService) CreatePhoto(userID int, request web.PhotoRequest) (web.CreatePhotoResponse, *exception.CustomError) {
	// validate request
	if err := s.Validate.Struct(request); err != nil {
		return web.CreatePhotoResponse{}, exception.ErrBadRequest(err.Error())
	}

	// create photo
	photo := domain.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   userID,
	}

	// insert photo to database
	savedPhoto, err := s.PhotoRepo.Create(photo)
	if err != nil {
		return web.CreatePhotoResponse{}, exception.ErrInternalServer("failed to save photo to database")
	}

	// return response
	response := web.CreatePhotoResponse{
		ID:        savedPhoto.ID,
		Title:     savedPhoto.Title,
		Caption:   savedPhoto.Caption,
		PhotoURL:  savedPhoto.PhotoURL,
		UserID:    savedPhoto.UserID,
		CreatedAt: savedPhoto.CreatedAt,
	}

	return response, nil
}

func (s *photoService) GetAllPhotos() ([]web.GetPhotoResponse, *exception.CustomError) {
	// get all photos from database
	photos, err := s.PhotoRepo.GetAll()
	if err != nil {
		return nil, exception.ErrInternalServer("failed to get photos from database")
	}

	var response []web.GetPhotoResponse
	for _, photo := range photos {
		// retrieve user data
		user, err := s.UserRepo.GetByID(photo.UserID)
		if err != nil {
			return nil, exception.ErrInternalServer("failed to get user data from database")
		}

		response = append(response, web.GetPhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: web.UserGetPhotoResponse{
				Username: user.Username,
				Email:    user.Email,
			},
		})
	}

	return response, nil
}

func (s *photoService) UpdatePhoto(userID, photoID int, request web.PhotoRequest) (web.UpdatePhotoResponse, *exception.CustomError) {
	// validate request
	if err := s.Validate.Struct(request); err != nil {
		return web.UpdatePhotoResponse{}, exception.ErrBadRequest(err.Error())
	}

	// get photo by id
	photo, err := s.PhotoRepo.GetByID(photoID)
	if err != nil {
		return web.UpdatePhotoResponse{}, exception.ErrNotFound("photo not found")
	}

	// check if the photo belongs to the user
	if photo.UserID != userID {
		return web.UpdatePhotoResponse{}, exception.ErrForbidden("you are not the owner of the photo")
	}

	// update photo
	photo.Title = request.Title
	photo.Caption = request.Caption
	photo.PhotoURL = request.PhotoURL

	// save photo to database
	updatedPhoto, err := s.PhotoRepo.Update(photo)
	if err != nil {
		return web.UpdatePhotoResponse{}, exception.ErrInternalServer("failed to update photo to database")
	}

	response := web.UpdatePhotoResponse{
		ID:        updatedPhoto.ID,
		Title:     updatedPhoto.Title,
		Caption:   updatedPhoto.Caption,
		PhotoURL:  updatedPhoto.PhotoURL,
		UserID:    updatedPhoto.UserID,
		UpdatedAt: updatedPhoto.UpdatedAt,
	}

	return response, nil
}

func (s *photoService) DeletePhoto(userID, photoID int) *exception.CustomError {
	// get photo by id
	photo, err := s.PhotoRepo.GetByID(photoID)
	if err != nil {
		return exception.ErrNotFound("photo not found")
	}

	// check if the photo belongs to the user
	if photo.UserID != userID {
		return exception.ErrForbidden("you are not the owner of the photo")
	}

	// delete photo
	if err := s.PhotoRepo.Delete(photo.ID); err != nil {
		return exception.ErrInternalServer("failed to delete photo")
	}

	return nil
}
