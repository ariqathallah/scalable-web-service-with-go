package service

import (
	"my-gram/exception"
	"my-gram/helper"
	"my-gram/model/domain"
	"my-gram/model/web"
	"my-gram/repository"
)

type SocialMediaService interface {
	Create(userID int, request web.SocialMediaRequest) (web.CreateSocialMediaResponse, *exception.CustomError)
	GetAllSocialMedias() ([]web.GetSocialMediaResponse, *exception.CustomError)
	Update(userID, socialMediaID int, request web.SocialMediaRequest) (web.UpdateSocialMediaResponse, *exception.CustomError)
	Delete(userID, socialMediaID int) *exception.CustomError
}

type socialMediaService struct {
	SocialMediaRepo repository.SocialMediaRepository
	UserRepo        repository.UserRepository
}

func NewSocialMediaService(socialMediaRepo repository.SocialMediaRepository, userRepo repository.UserRepository) *socialMediaService {
	return &socialMediaService{
		SocialMediaRepo: socialMediaRepo,
		UserRepo:        userRepo,
	}
}

func (s *socialMediaService) Create(userID int, request web.SocialMediaRequest) (web.CreateSocialMediaResponse, *exception.CustomError) {
	// validate request
	validateMessages := helper.SocialMediaValidate(request)
	for _, message := range validateMessages {
		return web.CreateSocialMediaResponse{}, exception.ErrBadRequest(message)
	}

	// create social media
	socialMedia := domain.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         userID,
	}

	// insert social media to database
	savedSocialMedia, err := s.SocialMediaRepo.Create(socialMedia)
	if err != nil {
		return web.CreateSocialMediaResponse{}, exception.ErrInternalServer("Failed to save social media to database")
	}

	// return response
	response := web.CreateSocialMediaResponse{
		ID:             savedSocialMedia.ID,
		Name:           savedSocialMedia.Name,
		SocialMediaURL: savedSocialMedia.SocialMediaURL,
		UserID:         savedSocialMedia.UserID,
		CreatedAt:      savedSocialMedia.CreatedAt,
	}

	return response, nil
}

func (s *socialMediaService) GetAllSocialMedias() ([]web.GetSocialMediaResponse, *exception.CustomError) {
	// get all social medias from database
	socialMedias, err := s.SocialMediaRepo.GetAll()
	if err != nil {
		return []web.GetSocialMediaResponse{}, exception.ErrInternalServer("Failed to get social medias from database")
	}

	var responses []web.GetSocialMediaResponse
	for _, socialMedia := range socialMedias {
		// retrieve user data
		user, err := s.UserRepo.GetByID(socialMedia.UserID)
		if err != nil {
			return nil, exception.ErrInternalServer("Failed to get user data from database")
		}

		response := web.GetSocialMediaResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt,
			UpdatedAt:      socialMedia.UpdatedAt,
			User: web.UserGetSocialMediaResponse{
				ID:              user.ID,
				Username:        user.Username,
				ProfileImageURL: "no profile image URL yet",
			},
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *socialMediaService) Update(userID, socialMediaID int, request web.SocialMediaRequest) (web.UpdateSocialMediaResponse, *exception.CustomError) {
	// validate request
	validateMessages := helper.SocialMediaValidate(request)
	for _, message := range validateMessages {
		return web.UpdateSocialMediaResponse{}, exception.ErrBadRequest(message)
	}

	// get social media by id
	socialMedia, err := s.SocialMediaRepo.GetByID(socialMediaID)
	if err != nil {
		return web.UpdateSocialMediaResponse{}, exception.ErrNotFound("Social media not found")
	}

	// check if the social media belongs to the user
	if socialMedia.UserID != userID {
		return web.UpdateSocialMediaResponse{}, exception.ErrForbidden("You are not allowed to update this social media")
	}

	// update social media
	socialMedia.Name = request.Name
	socialMedia.SocialMediaURL = request.SocialMediaURL

	// save social media to database
	updatedSocialMedia, err := s.SocialMediaRepo.Update(socialMedia)
	if err != nil {
		return web.UpdateSocialMediaResponse{}, exception.ErrInternalServer("Failed to update social media in database")
	}

	response := web.UpdateSocialMediaResponse{
		ID:             updatedSocialMedia.ID,
		Name:           updatedSocialMedia.Name,
		SocialMediaURL: updatedSocialMedia.SocialMediaURL,
		UserID:         updatedSocialMedia.UserID,
		UpdatedAt:      updatedSocialMedia.UpdatedAt,
	}

	return response, nil
}

func (s *socialMediaService) Delete(userID, socialMediaID int) *exception.CustomError {
	// get social media by ID
	socialMedia, err := s.SocialMediaRepo.GetByID(socialMediaID)
	if err != nil {
		return exception.ErrNotFound("Social media not found")
	}

	// check if the social media belongs to the user
	if socialMedia.UserID != userID {
		return exception.ErrForbidden("You are not allowed to delete this social media")
	}

	// delete social media
	if err := s.SocialMediaRepo.Delete(socialMediaID); err != nil {
		return exception.ErrInternalServer("Failed to delete social media from database")
	}

	return nil
}
