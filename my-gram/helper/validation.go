package helper

import (
	"fmt"
	"my-gram/model/web"

	"github.com/go-playground/validator"
)

func RegisterValidate(newUser web.RegisterRequest) []string {
	var validate = validator.New()
	err := validate.Struct(newUser)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMessage []string
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMessage = append(errorMessage, fmt.Sprintf("%s is required", e.Field()))
			case "min":
				errorMessage = append(errorMessage, fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param()))
			case "email":
				errorMessage = append(errorMessage, fmt.Sprintf("%s format is invalid", e.Field()))
			case "gte":
				errorMessage = append(errorMessage, fmt.Sprintf("%s must be at least %s years old", e.Field(), e.Param()))
			}
		}
		return errorMessage
	}
	return nil
}

func LoginValidate(user web.LoginRequest) []string {
	var validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMessage []string
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMessage = append(errorMessage, fmt.Sprintf("%s is required", e.Field()))
			case "email":
				errorMessage = append(errorMessage, fmt.Sprintf("%s format is invalid", e.Field()))
			}
		}
		return errorMessage
	}
	return nil
}

func UpdateUserValidate(user web.UpdateUserRequest) []string {
	var validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMessage []string
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMessage = append(errorMessage, fmt.Sprintf("%s is required", e.Field()))
			case "email":
				errorMessage = append(errorMessage, fmt.Sprintf("%s format is invalid", e.Field()))
			}
		}
		return errorMessage
	}
	return nil
}

func PhotoValidate(photo web.PhotoRequest) []string {
	var validate = validator.New()
	err := validate.Struct(photo)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMessage []string
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMessage = append(errorMessage, fmt.Sprintf("%s is required", e.Field()))
			}
		}
		return errorMessage
	}
	return nil
}

func CreateCommentValidate(comment web.CreateCommentRequest) []string {
	var validate = validator.New()
	err := validate.Struct(comment)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMessage []string
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMessage = append(errorMessage, fmt.Sprintf("%s is required", e.Field()))
			}
		}
		return errorMessage
	}
	return nil
}

func UpdateCommentValidate(comment web.UpdateCommentRequest) []string {
	var validate = validator.New()
	err := validate.Struct(comment)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMessage []string
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMessage = append(errorMessage, fmt.Sprintf("%s is required", e.Field()))
			}
		}
		return errorMessage
	}
	return nil
}

func SocialMediaValidate(socialMedia web.SocialMediaRequest) []string {
	var validate = validator.New()
	err := validate.Struct(socialMedia)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errorMessage []string
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMessage = append(errorMessage, fmt.Sprintf("%s is required", e.Field()))
			}
		}
		return errorMessage
	}
	return nil
}
