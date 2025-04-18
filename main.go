package main

import (
	"app/config"
	"app/internal/container"
	server "app/internal/grpc"
	"app/internal/middleware"
	"app/internal/router"
	"fmt"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	env := os.Getenv("ENV")
	r := gin.New()
	var logger *zap.Logger
	if env == "production" {
		gin.SetMode(gin.ReleaseMode) // 生产环境
		logger, _ = zap.NewProduction()

	} else {
		gin.SetMode(gin.DebugMode) // 开发环境
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()
	// 追溯Id
	r.Use(middleware.Trace)
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(middleware.RecoveryWithZap(logger))
	middleLog := middleware.NewLogger(logger)
	// 日志
	r.Use(middleLog.Logger)
	envConfig, envErr := config.InitConfig()
	if envErr != nil {
		logger.Error("配置错误", zap.String("traceId", envErr.Error()))
	}
	fmt.Println(envConfig.Service, config.Cfg, "envConfig")

	// 初始化grpc服务
	deps := container.InitContainer(logger)
	go server.IntServer(deps)
	router.RegisterRoutes(r, deps)

	// 运行服务器
	err := r.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
