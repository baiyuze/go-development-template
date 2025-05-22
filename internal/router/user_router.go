package router

import (
	"app/internal/handler"
	"app/internal/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterUserRoutes 注册所有路由
func RegisterUserRoutes(r *gin.Engine, container *dig.Container) {

	router := r.Group("user")
	err := container.Invoke(func(userHandler *handler.UserHandler, rpcHandler *handler.RpcHandler) {
		// 登录
		router.POST("/login", middleware.Jwt(false), userHandler.Login)
		//注册
		router.POST("/register", middleware.Jwt(false), userHandler.Register)
		//获取列表
		router.GET("/list", middleware.Jwt(true), userHandler.List)
		//jwt认证测试
		router.GET("/auth", middleware.Jwt(true), userHandler.TestAuth)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
