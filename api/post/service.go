package post

import (
	"github.com/igordonatti/REST-API/api/post/dto"
	"github.com/igordonatti/REST-API/api/post/model"
	"gorm.io/gorm"
)

type PostService struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{DB: db}
}

func (s *PostService) CreatePost(data dto.CreatePostDTO) (*model.Post, error) {
	post := model.Post{
		Text:   data.Text,
		BookID: uint(data.BookID),
		UserID: uint(data.UserID),
	}

	if err := s.DB.Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}
