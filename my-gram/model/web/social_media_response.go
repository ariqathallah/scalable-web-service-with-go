package web

import "time"

type CreateSocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type GetSocialMediaResponse struct {
	ID             int                        `json:"id"`
	Name           string                     `json:"name"`
	SocialMediaURL string                     `json:"social_media_url"`
	UserID         int                        `json:"user_id"`
	CreatedAt      time.Time                  `json:"created_at"`
	UpdatedAt      time.Time                  `json:"updated_at"`
	User           UserGetSocialMediaResponse `json:"user"`
}

type UpdateSocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
