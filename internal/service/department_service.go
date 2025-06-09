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
	Tree(context *gin.Context) ([]*model.Department, error)
	Update(c *gin.Context, id int, body *dto.DepartmentBody) error
	UpdateUsers(c *gin.Context, id int, body *dto.UsersIds) error
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
	var count int64
	var parentId = 0
	if err := s.db.Model(&model.Department{}).Count(&count).Error; err != nil {
		return err
	}

	if body.ParentId != nil {
		parentId = *body.ParentId
	}

	if err := s.db.Create(&model.Department{
		Name:        body.Name,
		Description: body.Description,
		Sort:        int(count + 1),
		ParentID:    parentId,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (s *departmentService) Delete(c *gin.Context, body *dto.DeleteIds) error {
	var departments []model.Department
	if err := s.db.Find(&departments, body.Ids).Error; err != nil {
		return err
	}

	if err := s.db.Model(&departments).Association("Users").Clear(); err != nil {
		return err
	}
	if err := s.db.Delete(&departments).Error; err != nil {
		return err
	}
	return nil
}

// Tree 列表
// 应该返回tree结构
func (s *departmentService) Tree(
	context *gin.Context,
) ([]*model.Department, error) {
	var tree []*model.Department
	var nodeMap = make(map[int]*model.Department)
	var departments []model.Department
	if err := s.db.Preload("Users", func(db *gorm.DB) *gorm.DB {
		return db.Select("users.name", "users.id", "users.email", "users.phone", "users.create_time", "users.update_time")
	}).Find(&departments).Error; err != nil {
		return nil, err
	}
	// 构造map
	for index, department := range departments {
		node := &departments[index]
		node.Children = []*model.Department{}
		nodeMap[department.ID] = node
	}
	//转换tree
	for index, department := range departments {
		var parentId = department.ParentID
		node := &departments[index]
		if parentId == 0 {
			tree = append(tree, node)
		} else {
			parentNode := nodeMap[parentId]
			parentNode.Children = append(parentNode.Children, node)

		}
	}

	return tree, nil
}
func (s *departmentService) Update(c *gin.Context, id int, body *dto.DepartmentBody) error {

	if err := s.db.Model(&model.Department{
		ID: id,
	}).Updates(&model.Department{
		Name:        body.Name,
		Description: body.Description,
		ParentID:    *body.ParentId,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (s *departmentService) UpdateUsers(c *gin.Context, id int, body *dto.UsersIds) error {
	var department model.Department
	var users []model.User
	if err := s.db.First(&department, id).Error; err != nil {
		return err
	}

	if err := s.db.Find(&users, body.Ids).Error; err != nil {
		return err
	}

	if err := s.db.Model(&department).Association("Users").Replace(&users); err != nil {
		return err
	}
	return nil
}
