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

type DictService interface {
	Create(c *gin.Context, body *dto.DepartmentBody) error
	Delete(c *gin.Context, body dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.Department]], error)
	Update(c *gin.Context, id int, body *dto.DepartmentBody) error
}

type dictService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewDictService(
	db *gorm.DB,
	log *log.LoggerWithContext) DictService {

	return &dictService{db: db, log: log}
}

func ProvideDictService(container *dig.Container) {
	if err := container.Provide(NewDictService); err != nil {
		panic(err)
	}
}

func (s *dictService) GetUserOne() (*model.User, error) {
	var user model.User
	if err := s.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Create 创建
func (s *dictService) Create(c *gin.Context, body *dto.DepartmentBody) error {

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

// List 获取所有的用户数据
func (s *dictService) List(c *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.Department]], error) {
	logger := s.log.WithContext(c)
	var roles []model.Department
	limit := query.PageSize
	offset := query.PageNum*query.PageSize - query.PageSize

	if result := s.db.
		Limit(limit).
		Offset(offset).
		Order("create_time asc").
		//不需要查询所有的角色用户
		//.Preload("Users", func(db *gorm.DB) *gorm.DB {
		//	return db.Select("users.name", "users.id", "users.email", "users.phone", "users.create_time", "users.update_time")
		//})
		Find(&roles); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[model.Department]](result.Error), result.Error
	}
	var count int64
	if result := s.db.Model(&model.Department{}).Count(&count); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[model.Department]](result.Error), result.Error
	}
	data := dto.ServiceSuccess(dto.List[model.Department]{
		Items:    roles,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    count,
	})
	return data, nil
}

// Delete 删除
func (s *dictService) Delete(c *gin.Context, body dto.DeleteIds) error {
	var roles []model.Role
	if err := s.db.Find(&roles, body.Ids).Error; err != nil {
		return err
	}
	// 清除权限关联
	if err := s.db.Model(&roles).Association("Permissions").Clear(); err != nil {
		return err
	}
	// 清除用户关联
	if err := s.db.Model(&roles).Association("Users").Clear(); err != nil {
		return err
	}
	if len(body.Ids) != 0 {
		s.db.Delete(&roles, body.Ids)
	}
	return nil
}

// updateRoleInfo 更新数据表字段
func updateInfo(db *gorm.DB, id int, body *dto.Role) error {
	if err := db.Model(&model.Role{}).Where("id = ?", id).Updates(&model.Role{
		Name:        body.Name,
		Description: body.Description,
	}).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新角色和关联关系，包括权限ID和用户ID
func (s *dictService) Update(c *gin.Context, id int, body *dto.DepartmentBody) error {

	return nil
}
