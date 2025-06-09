package router

import (
	"app/internal/handler"
	"app/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterDepartmentRoutes 注册所有路由
func RegisterDepartmentRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("departments")
	err := container.Invoke(func(departmentsHandler *handler.DepartmentHandler) {
		// 列表
		router.GET("/", middleware.Jwt(true), departmentsHandler.List)
		// 创建
		router.POST("/", middleware.Jwt(true), departmentsHandler.Create)
		// 删除
		router.DELETE("/", middleware.Jwt(true), departmentsHandler.Delete)
		// 修改
		router.PUT("/:id", middleware.Jwt(true), departmentsHandler.Update)
		// 绑定用户
		router.POST("/:id/users", middleware.Jwt(true), departmentsHandler.BindUser)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
