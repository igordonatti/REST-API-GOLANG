package dto

type CreatePostDTO struct {
	Text   string `json:"text" validate:"required"`
	BookID int    `json:"BookID" validate:"required"`
	UserID int    `json:"UserID" validate:"required"`
}
