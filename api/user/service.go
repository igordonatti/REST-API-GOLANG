package user

import (
	"github.com/igordonatti/REST-API/api/user/dto"
	"github.com/igordonatti/REST-API/api/user/model"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(data dto.CreateUserDTO) (*model.User, error) {
	user := model.User{
		Name:  data.Name,
		Email: data.Email,
	}
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
