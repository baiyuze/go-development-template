package router

import (
	_ "app/docs"
	"app/internal/dto"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
	"net/http"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine, container *dig.Container) {
	route := r.Group("api")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		// 可统一带上 traceId 日志
		c.JSON(http.StatusNotFound, dto.Fail(http.StatusNotFound, c.Request.RequestURI+"：接口不存在"))
	})
	RegisterUserRoutes(route, container)
	RegisterRpcRoutes(route, container)
	RegisterRolesRoutes(route, container)
	RegisterPermissionsRoutes(route, container)
	RegisterDepartmentRoutes(route, container)
	RegisterDictRoutes(route, container)

}
