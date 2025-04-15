package router

import (
	// "app/internal/controllers"
	// "app/internal/router"

	"app/internal/container"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine, deps *container.AppDependency) {
	RegisterUserRoutes(r, deps)

}
