package model

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Stars  string
	Text   string
	BookID uint // Chave estrangeira para Livro
	UserID uint // Chave estrangeira para Usuario
}
