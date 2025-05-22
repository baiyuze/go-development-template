package router

import (
	// "app/internal/controllers"
	// "app/internal/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine, container *dig.Container) {

	RegisterUserRoutes(r, container)
	RegisterRpcRoutes(r, container)
}
