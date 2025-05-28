package main

import (
	_ "app/docs"
	"app/internal/common/log"
	"app/internal/di"
	server "app/internal/grpc"
	"app/internal/middleware"
	"app/internal/router"
	"fmt"
	"go.uber.org/zap"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("ENV")
	isProduction := env == "production"
	r := gin.New()

	logger, _ := log.InitLogger()

	if isProduction {
		gin.SetMode(gin.ReleaseMode) // 生产环境
	} else {
		gin.SetMode(gin.DebugMode) // 开发环境
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error("failed to sync logger", zap.Error(err))
		}
	}(logger)
	// recover恢复
	r.Use(middleware.RecoveryWithZap(logger))
	middleLog := middleware.NewLogger(logger)

	// 追溯Id
	r.Use(middleware.Trace)
	// 日志
	r.Use(middleLog.Logger)

	// 白名单，暂时无用
	r.Use(middleware.AuthWhiteList)

	container := di.NewContainer()
	// 初始化grpc服务
	go server.IntServer(container)

	router.RegisterRoutes(r, container)
	// 运行服务器
	err := r.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
