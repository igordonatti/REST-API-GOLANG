package model

import (
	postModel "github.com/igordonatti/REST-API/api/post/model"
	ratingModel "github.com/igordonatti/REST-API/api/rating/model"
	"gorm.io/gorm"
)

type User struct {
	// visit https://gorm.io/docs/models.html#:~:text=on%20conventions.-,gorm.Model,-GORM%20provides%20a
	gorm.Model

	Name     string
	Email    string
	Password string
	PicUrl   string // colocar default null
	Autor    bool
	// Relacionamento 1:N com Postagem
	Posts []postModel.Post `gorm:"foreignKey:UserID"`
	// Relacionamento 1:N com Avaliacao
	Ratings []ratingModel.Rating `gorm:"foreignKey:UserID"`
}
