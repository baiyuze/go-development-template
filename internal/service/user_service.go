package service

import (
	"app/internal/models"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserOne() (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) GetUserOne() (*models.User, error) {
	var user models.User
	if err := s.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
