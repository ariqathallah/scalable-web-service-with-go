package web

import "time"

type CreatePhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetPhotoResponse struct {
	ID        int                  `json:"id"`
	Title     string               `json:"title"`
	Caption   string               `json:"caption"`
	PhotoURL  string               `json:"photo_url"`
	UserID    int                  `json:"user_id"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	User      UserGetPhotoResponse `json:"user"`
}

type UpdatePhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoGetCommentResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}
