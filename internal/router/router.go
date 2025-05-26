package router

import (
	// "app/internal/controllers"
	// "app/internal/router"

	"app/internal/dto"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine, container *dig.Container) {
	route := r.Group("api")
	r.NoRoute(func(c *gin.Context) {
		// 可统一带上 traceId 日志
		c.JSON(404, dto.Fail(http.StatusNotFound, c.Request.RequestURI+"：接口不存在"))
	})
	RegisterUserRoutes(route, container)
	RegisterRpcRoutes(route, container)
	RegisterRolesRoutes(route, container)
}
