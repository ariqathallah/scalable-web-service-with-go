package web

import "time"

type CreateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetCommentResponse struct {
	ID        int                     `json:"id"`
	Message   string                  `json:"message"`
	PhotoID   int                     `json:"photo_id"`
	UserID    int                     `json:"user_id"`
	UpdatedAt time.Time               `json:"updated_at"`
	CreatedAt time.Time               `json:"created_at"`
	User      UserGetCommentResponse  `json:"user"`
	Photo     PhotoGetCommentResponse `json:"photo"`
}

type UpdateCommentResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
