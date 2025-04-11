package router

import (
	// "app/internal/controllers"
	// "app/internal/router"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine, deps *AppDependency) {
	RegisterUserRoutes(r, deps, "user")

}
