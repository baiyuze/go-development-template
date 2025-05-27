package service

import (
	"app/internal/common/log"
	"app/internal/dto"
	"app/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"time"
)

type RolesService interface {
	Create(c *gin.Context, body model.Role) error
	Delete(c *gin.Context, body dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.Role]], error)
	UpdateRole(c *gin.Context, body dto.UserRoleRequest) error
	Update(c *gin.Context, id int, body *dto.Role) error
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

// UpdateRole 更新角色信息，包括权限ID和用户ID
func (s *rolesService) UpdateRole(c *gin.Context, body dto.UserRoleRequest) error {

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
func (s *rolesService) List(c *gin.Context, query dto.ListQuery) (dto.Result[dto.List[model.Role]], error) {
	logger := s.log.WithContext(c)
	var roles []model.Role
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
		return dto.ServiceFail[dto.List[model.Role]](result.Error), result.Error
	}
	var count int64
	if result := s.db.Model(&model.Role{}).Count(&count); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[model.Role]](result.Error), result.Error
	}
	data := dto.ServiceSuccess(dto.List[model.Role]{
		Items:    roles,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    count,
	})
	return data, nil
}

// Delete 删除
func (s *rolesService) Delete(c *gin.Context, body dto.DeleteIds) error {
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

// Update 更新角色和关联关系，包括权限ID和用户ID
func (s *rolesService) Update(c *gin.Context, id int, body *dto.Role) error {
	logger := s.log.WithContext(c)

	if len(body.Users) == 0 && len(body.Permissions) == 0 {
		//	只更新数据字段
		if err := s.db.Model(&model.Role{}).Where("id = ?", id).Updates(&model.Role{
			Name:        body.Name,
			Description: body.Description,
		}).Error; err != nil {
			return err
		}
	} else {
		var role model.Role
		if err := s.db.First(&role, id).Error; err != nil {
			return err
		}
		fmt.Println(body.Users, role, "body.Users--->")

		err := s.db.Transaction(func(tx *gorm.DB) error {
			// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
			//if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
			//	// 返回任何错误都会回滚事务
			//	return err
			//}
			//
			//if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
			//	return err
			//}

			//更新字段
			if len(body.Users) != 0 {
				//	更新依赖关系
				if err := tx.Model(&role).Association("Users").Replace(body.Users); err != nil {
					return err
				}
			}
			if len(body.Permissions) != 0 {
				//	更新依赖关系
				if err := tx.Model(&role).Association("Permissions").Replace(body.Users); err != nil {
					return err
				}
			}
			logger.Info("更新角色成功")
			return nil
		})
		if err != nil {
			return err
		}

	}

	return nil
}
