package router

import (
	"app/internal/handler"
	"app/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterPermissionsRoutes 注册所有路由
func RegisterPermissionsRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("permissions")
	err := container.Invoke(func(permissionsHandler *handler.PermissionsHandler) {
		// 列表
		router.GET("/", middleware.Jwt(true), permissionsHandler.List)
		// 创建
		router.POST("/", middleware.Jwt(true), permissionsHandler.Create)
		// 删除
		router.DELETE("/", middleware.Jwt(true), permissionsHandler.Delete)
		// 修改
		router.PUT("/:id", middleware.Jwt(true), permissionsHandler.Update)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
