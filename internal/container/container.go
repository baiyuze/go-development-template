// userHandler := controllers.NewUserHandler(DB)
package container

import (
	AppContext "app/internal/app_ontext"
	"app/internal/repo"
	"app/internal/service"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppDependency struct {
	DB          *gorm.DB
	Context     *AppContext.AppContext
	UserService service.UserService
	Logger      *zap.Logger
}

var Deps *AppDependency

// 初始化service
func InitContainer(logger *zap.Logger) *AppDependency {
	DB := repo.InitDB()
	// 初始化grpc客户端

	userService := service.NewUserService(DB)
	Deps = &AppDependency{
		UserService: userService,
		Logger:      logger,
	}
	return Deps
}

// 初始化GRPC 客户端
func InitClient(logger *zap.Logger) *AppDependency {
	DB := repo.InitDB()
	// 初始化grpc客户端

	ctx := AppContext.InitClient(logger)
	userService := service.NewUserService(DB)
	Deps = &AppDependency{
		UserService: userService,
		Context:     ctx,
		Logger:      logger,
	}
	return Deps
}
