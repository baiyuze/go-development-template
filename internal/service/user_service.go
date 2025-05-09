package service

import (
	"app/internal/model"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserOne() (*model.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func ProvideUserService(container *dig.Container) {
	container.Provide(NewUserService)
}

func (s *userService) GetUserOne() (*model.User, error) {
	var user model.User
	if err := s.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
