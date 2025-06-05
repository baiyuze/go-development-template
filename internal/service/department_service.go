package service

import (
	"app/internal/common/log"
	"app/internal/dto"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type DepartmentService interface {
	Create(c *gin.Context, body *dto.DepartmentBody) error
	Delete(c *gin.Context, body *dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.Permission]], error)
	Update(c *gin.Context, id int, body *dto.ReqPermissions) error
}

type departmentService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewDepartmentService(
	db *gorm.DB,
	log *log.LoggerWithContext) DepartmentService {

	return &departmentService{db: db, log: log}
}

func ProvideDepartmentService(container *dig.Container) {
	if err := container.Provide(NewDepartmentService); err != nil {
		panic(err)
	}
}

func (s *departmentService) Create(c *gin.Context, body *dto.DepartmentBody) error {
	if err := s.db.Create(&model.Permission{
		Name:        body.Name,
		Description: body.Description,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (s *departmentService) Delete(c *gin.Context, body *dto.DeleteIds) error {
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
func (s *departmentService) List(
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
func (s *departmentService) Update(c *gin.Context, id int, body *dto.ReqPermissions) error {
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
