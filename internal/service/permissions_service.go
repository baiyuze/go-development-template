package service

import (
	"app/internal/common/log"
	"app/internal/dto"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type PermissionsService interface {
	Create(c *gin.Context, body *dto.ReqPermissions) error
	Delete(c *gin.Context, body *dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.Permission]], error)
	Update(c *gin.Context, id int, body *dto.ReqPermissions) error
}

type permissionsService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewPermissionsService(
	db *gorm.DB,
	log *log.LoggerWithContext) PermissionsService {

	return &permissionsService{db: db, log: log}
}

func ProvidePermissionsService(container *dig.Container) {
	if err := container.Provide(NewPermissionsService); err != nil {
		panic(err)
	}
}

func (s *permissionsService) Create(c *gin.Context, body *dto.ReqPermissions) error {
	if err := s.db.Create(&model.Permission{
		Name:        body.Name,
		Description: body.Description,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (s *permissionsService) Delete(c *gin.Context, body *dto.DeleteIds) error {
	//先清空关联关系，然后再删除
	// 查找权限列表
	var permissions []model.Permission
	if err := s.db.Find(&permissions, body.Ids).Error; err != nil {
		return err
	}

	if err := s.db.Model(&permissions).Association("Roles").Clear(); err != nil {
		return err
	}
	if err := s.db.Delete(&model.Permission{}, body.Ids).Error; err != nil {
		return err
	}
	return nil
}

// List 列表
func (s *permissionsService) List(
	context *gin.Context,
	query dto.ListQuery,
) (dto.Result[dto.List[model.Permission]], error) {
	var permissions []model.Permission
	limit := query.PageSize
	offset := query.PageNum*query.PageSize - query.PageSize
	if err := s.db.
		Limit(limit).
		Offset(offset).
		Find(&permissions).Error; err != nil {
		return dto.ServiceFail[dto.List[model.Permission]](nil), err
	}
	return dto.ServiceSuccess(dto.List[model.Permission]{
		Items:    permissions,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    1,
	}), nil
}
func (s *permissionsService) Update(c *gin.Context, id int, body *dto.ReqPermissions) error {
	if err := s.db.Model(&model.Permission{
		ID: id,
	}).Updates(&model.Permission{
		Name:        body.Name,
		Description: body.Description,
	}).Error; err != nil {
		return err
	}
	return nil
}
