package web

type CreateCommentRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoID int    `json:"photo_id" validate:"required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"required"`
}
