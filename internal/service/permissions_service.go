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
	Create(c *gin.Context, body model.Role) error
	Delete(c *gin.Context, body dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.Role]], error)
	Update(c *gin.Context, id int, body *dto.Role) error
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

func (s *permissionsService) Create(c *gin.Context, body model.Role) error {
	return nil
}

func (s *permissionsService) Delete(c *gin.Context, body dto.DeleteIds) error {
	return nil
}

func (s *permissionsService) List(context *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.Role]], error) {
	return dto.ServiceSuccess(dto.List[model.Role]{
		Items:    []model.Role{},
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    1,
	}), nil
}
func (s *permissionsService) Update(c *gin.Context, id int, body *dto.Role) error {
	return nil
}
