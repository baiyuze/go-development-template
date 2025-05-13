package di

import (
	"app/config"
	"app/internal/common/logx"
	grpcContainer "app/internal/grpc/container"
	"app/internal/handler"
	"app/internal/repo"
	"app/internal/service"

	"go.uber.org/dig"
	"go.uber.org/zap"
)

// NewContainer 创建并初始化 DI 容器
func NewContainer(logger *zap.Logger) *dig.Container {
	// 注册各模块的依赖
	container := dig.New()
	// 注入logger
	container.Provide(func() *zap.Logger {
		return logger
	})
	// 公共日志管理器
	logx.NewProvideLogger(container)
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
