package router

import (
	"app/internal/handler"
	"app/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterRolesRoutes 注册所有路由
func RegisterRolesRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("roles")
	err := container.Invoke(func(rolesHandler *handler.RolesHandler) {
		// 修改角色
		router.POST("/set-role", middleware.Jwt(true), rolesHandler.UpdateRole)
		// 列表
		router.GET("/", middleware.Jwt(true), rolesHandler.List)
		// 创建
		router.POST("/", middleware.Jwt(true), rolesHandler.Create)
		// 删除
		router.DELETE("/", middleware.Jwt(true), rolesHandler.Delete)
		// 修改
		router.PUT("/:id", middleware.Jwt(true), rolesHandler.Update)

	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
