package router

import (
	"app/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
// internal/router/user_router.go
func RegisterUserRoutes(r *gin.Engine, deps *AppDependency, groupPath string) {

	router := r.Group(groupPath)
	userHandler := handler.NewUserHandler(deps.UserService)

	router.GET("/", userHandler.HomeHandler)
	// r.GET("/user/:id", handler.GetUser)
}
