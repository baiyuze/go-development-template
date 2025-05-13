package main

import (
	"app/internal/di"
	server "app/internal/grpc"
	"app/internal/middleware"
	"app/internal/router"
	"fmt"
	"os"

	"app/internal/common/logx"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	env := os.Getenv("ENV")
	isProduction := env == "production"
	r := gin.New()

	var logger *zap.Logger
	if isProduction {
		gin.SetMode(gin.ReleaseMode) // 生产环境
		logger, _ = logx.InitLogger()
	} else {
		gin.SetMode(gin.DebugMode) // 开发环境
		logger, _ = logx.InitLogger()
	}
	defer logger.Sync()
	// recover恢复
	r.Use(middleware.RecoveryWithZap(logger))
	middleLog := middleware.NewLogger(logger)

	// 追溯Id
	r.Use(middleware.Trace)
	// 认证白名单
	r.Use(middleware.AuthWhiteList)
	r.Use(middleware.Jwt)

	// 日志
	r.Use(middleLog.Logger)
	// envConfig, envErr := config.InitConfig()
	// if envErr != nil {
	// 	logger.Error("配置错误", zap.String("traceId", envErr.Error()))
	// }

	// fmt.Println(envConfig.Service, config.Cfg, "envConfig")
	// Deps := container.InitContainer(logger)

	container := di.NewContainer(logger)

	go func() {
		go server.IntServer(container)
	}()

	// AppContext.InitClient(logger)

	// 初始化grpc服务
	router.RegisterRoutes(r, container)

	// 运行服务器
	err := r.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
