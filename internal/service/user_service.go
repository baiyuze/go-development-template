package service

import (
	errs "app/internal/common/error"
	"app/internal/common/jwt"
	"app/internal/dto"
	"app/internal/model"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserOne() (*model.User, error)
	Login(body dto.LoginBody) (string, error)
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

// 登录进行校验返回token
func (s *userService) Login(body dto.LoginBody) (string, error) {
	var user model.User

	result := s.db.Where("account = ?", body.Account).First(&user)
	if result.Error != nil {
		return "", result.Error
	}
	psd := sha256.Sum256([]byte(body.Password))
	hashPsd := hex.EncodeToString(psd[:])
	if user.Account == body.Account && hashPsd == user.Password {
		// 调用jwt
		sign, err := jwt.Auth(user, time.Now().Unix()+1000*60*60*2)
		if err != nil {
			errs.MustNoErr(err, "token创建失败")
			return "", err
		}
		fmt.Println(sign, err, "vuserInfouserInfouserInfouserInfo_----->")
	}
	return "", nil
	// return
}
