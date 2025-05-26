package service

import (
	"app/internal/common/log"
	"app/internal/dto"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"time"
)

type RolesService interface {
	Create(c *gin.Context, body model.Role) error
	Delete(c *gin.Context, body dto.RegBody) error
	List(context *gin.Context, query dto.ListQuery) (dto.Result[dto.List[dto.UserWithRole]], error)
	Update(c *gin.Context, body dto.UserRoleRequest) error
}

type rolesService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewRolesService(
	db *gorm.DB,
	log *log.LoggerWithContext) RolesService {

	return &rolesService{db: db, log: log}
}

func ProvideRolesService(container *dig.Container) {
	if err := container.Provide(NewRolesService); err != nil {
		panic(err)
	}
}

func (s *rolesService) GetUserOne() (*model.User, error) {
	var user model.User
	if err := s.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Create 创建
func (s *rolesService) Create(c *gin.Context, body model.Role) error {

	//logger := s.log.WithContext(c)
	result := s.db.Create(&model.Role{
		Name:        body.Name,
		Description: body.Description,
		CreateTime:  time.Now(),
	})
	
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Update 更新
func (s *rolesService) Update(c *gin.Context, body dto.UserRoleRequest) error {

	var user model.User
	if err := s.db.First(&user, body.ID).Error; err != nil {
		return err
	}
	//先查出来用户，再查出来角色对象，然后通过用户去更新替换角色id
	// 查出要绑定的角色对象
	var roles []model.Role
	if err := s.db.Where("id IN ?", body.RoleIds).Find(&roles).Error; err != nil {
		return err
	}

	if err := s.db.Model(&user).Association("Roles").Replace(&roles); err != nil {
		return err
	}
	return nil
}

// List 获取所有的用户数据
func (s *rolesService) List(c *gin.Context, query dto.ListQuery) (dto.Result[dto.List[dto.UserWithRole]], error) {
	logger := s.log.WithContext(c)
	var users []dto.UserWithRole
	limit := query.PageSize
	offset := query.PageNum*query.PageSize - query.PageSize

	if result := s.db.
		Table("users").
		Select(
			"users.id, users.name, users.account, users.create_time, users.update_time").
		//", roles.name as role_name, user.role_id")
		//Joins("LEFT JOIN roles ON user.role_id = roles.id").
		Limit(limit).
		Offset(offset).
		Order("create_time asc").
		Scan(&users); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[dto.UserWithRole]](result.Error), result.Error
	}
	var count int64
	if result := s.db.Model(&model.User{}).Count(&count); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[dto.UserWithRole]](result.Error), result.Error
	}
	data := dto.ServiceSuccess(dto.List[dto.UserWithRole]{
		Items:    users,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    count,
	})
	return data, nil
}

func (s *rolesService) Delete(c *gin.Context, body dto.RegBody) error {
	return nil
}
