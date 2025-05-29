package handler

import (
	"app/internal/common/log"
	"app/internal/grpc/container"
	"app/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type PermissionsHandler struct {
	service *service.PermissionsService
	log     *log.LoggerWithContext
	client  *container.Clients
}

func NewPermissionsHandler(s *service.PermissionsService, l *log.LoggerWithContext, client *container.Clients) *PermissionsHandler {
	return &PermissionsHandler{
		service: s,
		log:     l,
		client:  client,
	}
}

func ProviderPermissionsHandler(container *dig.Container) {
	if err := container.Provide(NewPermissionsHandler); err != nil {
		return
	}
}

// Create 创建
// @Summary 创建
// @Tags 权限码模块
// @Accept  json
// @Params data body model.Permission
// @Success 200  {object} dto.Response[any]
// @Router /api/permissions [post]
func (h *PermissionsHandler) Create(c *gin.Context) {

}

// Update 更新
// @Summary 更新
// @Tags 权限码模块
// @Accept  json
// @Params data body model.Permission
// @Success 200  {object} dto.Response[any]
// @Router /api/permissions [put]
func (h *PermissionsHandler) Update(c *gin.Context) {

}

// List 查询
// @Summary 查询
// @Tags 权限码模块
// @Accept  json
// @Params data body model.Permission
// @Success 200  {object} dto.Response[any]
// @Router /api/permissions [get]
func (h *PermissionsHandler) List(c *gin.Context) {

}

// Delete 删除
// @Summary 删除
// @Tags 权限码模块
// @Accept  json
// @Params data body model.Permission
// @Success 200  {object} dto.Response[any]
// @Router /api/permissions [get]
func (h *PermissionsHandler) Delete(c *gin.Context) {

}
