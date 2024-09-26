package book

import "gorm.io/gorm"

type BookService struct {
	DB *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{DB: db}
}
