package router

import (
	"app/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
// internal/router/user_router.go
func RegisterUserRoutes(r *gin.Engine, deps *AppDependency) {

	router := r.Group("user")
	userHandler := handler.NewUserHandler(deps.UserService)

	router.GET("/", userHandler.HomeHandler)
	// 测试RPC
	router.GET("/test", userHandler.TestRpc)
}
