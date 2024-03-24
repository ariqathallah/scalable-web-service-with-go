package web

type SocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required"`
}
