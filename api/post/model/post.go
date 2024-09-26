package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Text          string
	DateOfPublish time.Time
	Status        string
	BookID        uint // chave estrangeira para Livro
	UserID        uint //chave estrangera para Usuario
}
