package service

import (
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
	Register(c *gin.Context, body dto.RegBody) error
	List(context *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.User]], error)
}

type userService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewUserService(
	db *gorm.DB,
	log *log.LoggerWithContext) UserService {

	return &userService{db: db, log: log}
}

func ProvideUserService(container *dig.Container) {
	if err := container.Provide(NewUserService); err != nil {
		panic(err)
	}
}

func (s *userService) GetUserOne() (*model.User, error) {
	var user model.User
	if err := s.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Login 登录进行校验返回token
func (s *userService) Login(c *gin.Context, body dto.LoginBody) dto.Result[dto.LoginResult] {
	var user model.User

	result := s.db.Where("account = ?", body.Account).First(&user)
	if result.Error != nil {
		return dto.ServiceFail[dto.LoginResult](errors.New("密码错误,请检查账号密码"))
	}
	psd := sha256.Sum256([]byte(body.Password))
	hashPsd := hex.EncodeToString(psd[:])
	if user.Account == body.Account && hashPsd == *user.Password {
		// 调用jwt
		//两小时过期
		sign, err := jwt.Auth(user, time.Now().Add(2*time.Hour).Unix())
		if err != nil {
			return dto.ServiceFail[dto.LoginResult](err)
		}
		//7天过期
		refreshToken, err := jwt.Auth(user,
			time.Now().Add(24*7*time.Hour).Unix())
		if err != nil {
			return dto.ServiceFail[dto.LoginResult](err)
		}

		return dto.ServiceSuccess(dto.LoginResult{
			Token:        sign,
			RefreshToken: refreshToken,
			UserInfo: &dto.UserInfo{
				Account: user.Account,
				Name:    user.Name,
				Id:      float64(user.ID),
			},
		})
	}
	return dto.ServiceFail[dto.LoginResult](errors.New("密码错误,请检查账号密码"))
}

// Register 注册
func (s *userService) Register(c *gin.Context, body dto.RegBody) error {

	logger := s.log.WithContext(c)
	var user model.User
	result := s.db.Where("account = ?", *body.Account).Find(&user)

	if result.Error == nil {
		if user.Account == *body.Account {
			return errors.New(*body.Account + "当前账号已经存在")
		}
	}
	psd := sha256.Sum256([]byte(*body.Password))
	hashPsd := hex.EncodeToString(psd[:])
	if body.Password != nil {
		result := s.db.Create(&model.User{
			Account:    *body.Account,
			Password:   &hashPsd,
			Name:       *body.Name,
			CreateTime: time.Now(),
		})
		if result.Error != nil {
			logger.Error(result.Error.Error())
			return result.Error
		}
	}
	return nil
}

func (s *userService) Update(c *gin.Context, body dto.RegBody) error { return nil }

// List 获取所有的用户数据
func (s *userService) List(c *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.User]], error) {
	logger := s.log.WithContext(c)
	users := make([]model.User, 0)
	limit := query.PageSize
	offset := query.PageNum*query.PageSize - query.PageSize

	if result := s.db.
		Select(
			"id",
			"name",
			"account",
			"create_time",
			"update_time",
			"role_id").Limit(limit).Offset(offset).Order("create_time asc").Find(&users); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[model.User]](result.Error), result.Error
	}
	var count int64
	if result := s.db.Model(&model.User{}).Count(&count); result.Error != nil {
		logger.Error(result.Error.Error())
	}
	data := dto.ServiceSuccess(dto.List[model.User]{
		Items:    users,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    count,
	})
	return data, nil
}
