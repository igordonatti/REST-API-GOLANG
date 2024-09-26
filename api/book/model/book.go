package model

import (
	"time"

	postModel "github.com/igordonatti/REST-API/api/post/model"
	ratingModel "github.com/igordonatti/REST-API/api/rating/model"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string    `validate:"required"`
	Author        string    `validate:"required"`
	ISBN          string    `validate:"required"`
	DateOfPublish time.Time `validate:"required"`
	Genre         string    `validate:"required"`
	Publisher     string    `validate:"required"`
	Synopsis      string    `validate:"required"`

	// Relacionamento 1:N com Postagem
	Posts []postModel.Post `gorm:"foreignKey:BookID"`
	// Relacionamento 1:N com Avaliacao
	Ratings []ratingModel.Rating `gorm:"foreignKey:BookID"`
}
