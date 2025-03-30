package main

import (
	"app/routes"
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
	// 注册路由
	routes.RegisterRoutes(r)

	// 运行服务器
	r.Run(":8800")
}
