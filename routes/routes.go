package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	prefixRoute := r.Group("/projectApi")
	prefixRoute.GET("/", controllers.HomeHandler)    // 主页
	prefixRoute.GET("/env", controllers.FileHandler) // 文件
	prefixRoute.POST("create", controllers.CreateWidget)
}
