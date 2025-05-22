package router

import (
	"app/internal/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterRoutes 注册所有路由
// internal/router/user_router.go
func RegisterUserRoutes(r *gin.Engine, container *dig.Container) {

	router := r.Group("user")
	err := container.Invoke(func(userHandler *handler.UserHandler, rpcHandler *handler.RpcHandler) {
		// 登录
		router.POST("/login", userHandler.Login)
		//注册
		router.POST("/register", userHandler.Register)
		//获取列表
		router.GET("/list", userHandler.List)

		router.GET("/auth", userHandler.TestAuth)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
