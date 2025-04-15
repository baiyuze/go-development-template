// userHandler := controllers.NewUserHandler(DB)
package container

import (
	AppContext "app/internal/app_ontext"
	"app/internal/repo"
	"app/internal/service"

	"gorm.io/gorm"
)

type AppDependency struct {
	DB          *gorm.DB
	Context     *AppContext.AppContext
	UserService service.UserService
}

var Deps *AppDependency

// 初始化service
func InitContainer() *AppDependency {
	DB := repo.InitDB()
	ctx := AppContext.InitClient()
	userService := service.NewUserService(DB)
	Deps = &AppDependency{
		UserService: userService,
		Context:     ctx,
	}
	return Deps
}
