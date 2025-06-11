package router

import (
	"app/internal/handler"
	"app/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterDictRoutes 注册所有路由
func RegisterDictRoutes(r *gin.RouterGroup, container *dig.Container) {

	router := r.Group("dicts")
	err := container.Invoke(func(dictHandler *handler.DictHandler) {
		// 列表
		router.GET("/", middleware.Jwt(true), dictHandler.List)
		// 创建
		router.POST("/", middleware.Jwt(true), dictHandler.Create)
		// 删除
		router.DELETE("/", middleware.Jwt(true), dictHandler.Delete)
		// 修改
		router.PUT("/:id", middleware.Jwt(true), dictHandler.Update)
		// 根据code获取options
		router.GET("/:code", middleware.Jwt(true), dictHandler.GetOptions)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
