package service

import (
	errs "app/internal/common/error"
	"app/internal/common/jwt"
	"app/internal/common/log"
	"app/internal/dto"
	"app/internal/model"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserOne() (*model.User, error)
	Login(c *gin.Context, body dto.LoginBody) dto.Result[dto.LoginResult]
}

type userService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewUserService(db *gorm.DB, log *log.LoggerWithContext) UserService {
	return &userService{db: db, log: log}
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
func (s *userService) Login(c *gin.Context, body dto.LoginBody) dto.Result[dto.LoginResult] {
	// logger := s.log.WithContext(c)
	var user model.User

	result := s.db.Where("account = ?", body.Account).First(&user)
	if result.Error != nil {
		return dto.Result[dto.LoginResult]{Data: dto.LoginResult{}, Error: errors.New("密码错误,请检查账号密码")}
	}
	psd := sha256.Sum256([]byte(body.Password))
	hashPsd := hex.EncodeToString(psd[:])
	if user.Account == body.Account && hashPsd == user.Password {
		// 调用jwt
		//两小时过期
		sign, err := jwt.Auth(user, time.Now().Add(2*time.Hour).Unix())
		if err != nil {
			errs.MustNoErr(err, "token创建失败")
			return dto.Result[dto.LoginResult]{Data: dto.LoginResult{}, Error: err}
		}
		//7天过期
		flushSign, err := jwt.Auth(user, time.Now().Add(24*7*time.Hour).Unix())
		if err != nil {
			errs.MustNoErr(err, "token创建失败")
			return dto.Result[dto.LoginResult]{Data: dto.LoginResult{}, Error: err}
		}
		return dto.Result[dto.LoginResult]{Data: dto.LoginResult{
			Token:      sign,
			FlushToken: flushSign,
			UserInfo: &dto.UserInfo{
				Account: user.Account,
				Name:    user.Name,
				Id:      float64(user.ID),
			},
		}, Error: nil}

	}
	return dto.Result[dto.LoginResult]{Data: dto.LoginResult{}, Error: errors.New("密码错误,请检查账号密码")}
}
