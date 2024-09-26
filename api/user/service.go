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
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		PicUrl:   data.PicUrl,
	}

	if err := s.DB.Where("email = ?", data.Email).First(&user).Error; err == nil {
		// Se o usuário já existe, retorna um erro específico
		return nil, err
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

func (s *UserService) FindUserByID(id uint) (*model.User, error) {
	var user model.User

	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) DeleteUserByID(id uint) (*model.User, error) {
	var user model.User

	// Verifica se o usuário com o ID fornecido existe no banco
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
