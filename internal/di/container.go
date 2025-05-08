package di

import (
	"app/config"
	grpcContainer "app/internal/grpc/container"
	"app/internal/handler"
	"app/internal/repo"
	"app/internal/service"

	"go.uber.org/dig"
)

// NewContainer 创建并初始化 DI 容器
func NewContainer() *dig.Container {
	// 注册各模块的依赖
	container := dig.New()
	// 获取客户端grpc
	grpcContainer.NewProvideClients(container)
	// 配置
	config.ProvideConfig(container)
	// 数据库
	repo.ProvideDB(container)
	// 用户服务
	service.ProvideUserService(container)
	// controller
	handler.ProviderUserHandler(container)

	return container
}
