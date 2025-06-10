package router

import (
	"app/internal/handler"
	"app/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterDicttRoutes 注册所有路由
func RegisterDicttRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("dict")
	err := container.Invoke(func(dictHandler *handler.DictHandler) {
		// 列表
		router.GET("/", middleware.Jwt(true), dictHandler.List)
		// 创建
		router.POST("/", middleware.Jwt(true), dictHandler.Create)
		// 删除
		router.DELETE("/", middleware.Jwt(true), dictHandler.Delete)
		// 修改
		router.PUT("/:id", middleware.Jwt(true), dictHandler.Update)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
