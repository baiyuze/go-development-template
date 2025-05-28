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
		// 列表
		router.GET("/", middleware.Jwt(true), rolesHandler.List)
		// 创建
		router.POST("/", middleware.Jwt(true), rolesHandler.Create)
		// 删除
		router.DELETE("/", middleware.Jwt(true), rolesHandler.Delete)
		// 修改
		router.PUT("/:id", middleware.Jwt(true), rolesHandler.Update)
		// 修改角色绑定的权限码
		router.PUT("/permissions/:id", middleware.Jwt(true), rolesHandler.UpdatePermissions)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
