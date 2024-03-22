package service

import (
	"my-gram/model/web"
	"my-gram/repository"

	"github.com/go-playground/validator"
)

type UserService interface {
	Login(request web.LoginRequest) (string, error)
	Register(request web.RegisterRequest) (string, error)
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

func (s *userService) Login(request web.LoginRequest) (string, error) {
	return "", nil
}

func (s *userService) Register(request web.RegisterRequest) (string, error) {
	return "", nil
}
