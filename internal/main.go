package main

import (
	"app/internal/container"
	"app/internal/router"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode) // 生产环境
	} else {
		gin.SetMode(gin.DebugMode) // 开发环境
	}
	r := gin.Default()
	deps := container.InitContainer()
	router.RegisterRoutes(r, deps)

	// 运行服务器
	r.Run(":8888")
}
