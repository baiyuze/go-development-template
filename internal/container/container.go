// userHandler := controllers.NewUserHandler(DB)
package container

import (
	"app/internal/repo"
	"app/internal/router"
	"app/internal/service"
)

// 初始化service
func InitContainer() *router.AppDependency {
	DB := repo.InitDB()

	userService := service.NewUserService(DB)
	deps := &router.AppDependency{
		UserService: userService,
	}
	return deps
}
