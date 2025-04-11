package router

import (
	"app/internal/service"

	"gorm.io/gorm"
)

type AppDependency struct {
	DB          *gorm.DB
	UserService service.UserService
}
