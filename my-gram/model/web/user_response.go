package web

import "time"

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UpdateUserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserGetPhotoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserGetCommentResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserGetSocialMediaResponse struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
}
