package main

import (
	"app/config"
	"app/internal/container"
	server "app/internal/grpc"
	"app/internal/router"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("ENV")
	r := gin.New()
	r.Use(gin.Logger(), gin.RecoveryWithWriter(os.Stderr))
	envConfig, envErr := config.InitConfig()
	if envErr != nil {
		fmt.Println("配置错误", envErr)
	}
	fmt.Println(envConfig.Service, config.Cfg, "envConfig")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode) // 生产环境
	} else {
		gin.SetMode(gin.DebugMode) // 开发环境
	}
	// 初始化grpc服务
	deps := container.InitContainer()
	go server.IntServer(deps)
	router.RegisterRoutes(r, deps)

	// 运行服务器
	err := r.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
