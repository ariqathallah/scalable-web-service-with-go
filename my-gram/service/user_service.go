package service

import (
	"my-gram/exception"
	"my-gram/helper"
	"my-gram/model/domain"
	"my-gram/model/web"
	"my-gram/repository"

	"github.com/go-playground/validator"
)

type UserService interface {
	Register(request web.RegisterRequest) (web.RegisterResponse, *exception.CustomError)
	Login(request web.LoginRequest) (web.LoginResponse, *exception.CustomError)
	UpdateUser(param, userID int, request web.UpdateUserRequest) (web.UpdateUserResponse, *exception.CustomError)
	DeleteUser(param, userID int) *exception.CustomError
}

type userService struct {
	Validate *validator.Validate
	UserRepo repository.UserRepository
}

func NewUserService(validate *validator.Validate, userRepo repository.UserRepository) *userService {
	return &userService{
		Validate: validate,
		UserRepo: userRepo,
	}
}

func (s *userService) Register(request web.RegisterRequest) (web.RegisterResponse, *exception.CustomError) {
	// validate request
	if err := s.Validate.Struct(request); err != nil {
		return web.RegisterResponse{}, exception.ErrBadRequest(err.Error())
	}

	// check if email already registered
	if _, err := s.UserRepo.GetByEmail(request.Email); err == nil {
		return web.RegisterResponse{}, exception.ErrBadRequest("email already registered")
	}

	// check if username already registered
	if _, err := s.UserRepo.GetByUsername(request.Username); err == nil {
		return web.RegisterResponse{}, exception.ErrBadRequest("username already registered")
	}

	// create user
	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Age:      request.Age,
	}
	user.SetPassword(request.Password)

	// insert user to database
	savedUser, err := s.UserRepo.Create(user)
	if err != nil {
		return web.RegisterResponse{}, exception.ErrInternalServer("failed to save user to database")
	}

	userResponse := web.RegisterResponse{
		ID:       savedUser.ID,
		Username: savedUser.Username,
		Email:    savedUser.Email,
		Age:      savedUser.Age,
	}

	return userResponse, nil
}

func (s *userService) Login(request web.LoginRequest) (web.LoginResponse, *exception.CustomError) {
	// validate request
	if err := s.Validate.Struct(request); err != nil {
		return web.LoginResponse{}, exception.ErrBadRequest("invalid request")
	}

	// get user by email
	user, err := s.UserRepo.GetByEmail(request.Email)
	if err != nil {
		return web.LoginResponse{}, exception.ErrBadRequest("invalid email or password")
	}

	// check if password is not match
	if err := user.CheckPassword(request.Password); err != nil {
		return web.LoginResponse{}, exception.ErrBadRequest("invalid email or password")
	}

	// generate jwt token
	token, err := helper.GenerateJWT(user.ID)
	if err != nil {
		return web.LoginResponse{}, exception.ErrInternalServer("failed to generate jwt token")
	}

	// return response
	response := web.LoginResponse{
		Token: token,
	}

	return response, nil
}

func (s *userService) UpdateUser(param, userID int, request web.UpdateUserRequest) (web.UpdateUserResponse, *exception.CustomError) {
	// validate request
	if err := s.Validate.Struct(request); err != nil {
		return web.UpdateUserResponse{}, exception.ErrBadRequest(err.Error())
	}

	// get user by id
	user, err := s.UserRepo.GetByID(param)
	if err != nil {
		return web.UpdateUserResponse{}, exception.ErrBadRequest("user not found")
	}

	// check if param is not match with userID
	if user.ID != userID {
		return web.UpdateUserResponse{}, exception.ErrForbidden("forbidden")
	}

	// check if email already registered
	if _, err := s.UserRepo.GetByEmail(request.Email); err == nil {
		return web.UpdateUserResponse{}, exception.ErrBadRequest("email already registered")
	}

	// check if username already registered
	if _, err := s.UserRepo.GetByUsername(request.Username); err == nil {
		return web.UpdateUserResponse{}, exception.ErrBadRequest("username already registered")
	}

	// update user
	user.Username = request.Username
	user.Email = request.Email

	// save user to database
	updatedUser, err := s.UserRepo.Update(user)
	if err != nil {
		return web.UpdateUserResponse{}, exception.ErrInternalServer("failed to save user to database")
	}

	// return response
	response := web.UpdateUserResponse{
		ID:        updatedUser.ID,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		Age:       updatedUser.Age,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}

func (s *userService) DeleteUser(param, userID int) *exception.CustomError {
	// get user by id
	user, err := s.UserRepo.GetByID(param)
	if err != nil {
		return exception.ErrBadRequest("user not found")
	}

	// check if param is not match with userID
	if user.ID != userID {
		return exception.ErrForbidden("forbidden")
	}

	// delete user
	if err := s.UserRepo.Delete(user.ID); err != nil {
		return exception.ErrInternalServer("failed to delete user")
	}

	return nil
}
