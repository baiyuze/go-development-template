package handler

import (
	errs "app/internal/common/error"
	log "app/internal/common/log"
	"app/internal/dto"
	"app/internal/grpc/container"
	"app/internal/model"
	"app/internal/service"
	"app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
)

type RolesHandler struct {
	service service.RolesService
	clients *container.Clients
	log     *log.LoggerWithContext
}

func NewRolesHandler(
	service service.RolesService,
	clients *container.Clients,
	log *log.LoggerWithContext,
) *RolesHandler {
	return &RolesHandler{
		service: service,
		clients: clients,
		log:     log,
	}
}

func ProviderRolesHandler(container *dig.Container) {
	err := container.Provide(NewRolesHandler)
	if err != nil {
		return
	}
}

// Create 注册
func (h *RolesHandler) Create(c *gin.Context) {
	logger := h.log.WithContext(c)

	var body model.Role
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err.Error())
		logger.Error(err.Error())
		return
	}

	if len(body.Name) != 0 && len(body.Description) != 0 {

		if err := h.service.Create(c, body); err != nil {
			errs.FailWithJSON(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, dto.Ok[any](nil))
		return
	} else {
		errs.FailWithJSON(c, "账号或密码不存在")
		return
	}

}
func (h *RolesHandler) List(c *gin.Context) {
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")

	result, err := h.service.List(c, utils.HandleQuery(pageNum, pageSize))
	if err != nil {
		errs.FailWithJSON(c, err.Error())
	} else {
		c.JSON(http.StatusOK, dto.Ok(result.Data))
	}
}

// UpdateRole 修改角色，设置角色
func (h *RolesHandler) UpdateRole(c *gin.Context) {
	logger := h.log.WithContext(c)
	var body dto.UserRoleRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		return
	}

	if len(body.RoleIds) == 0 || body.ID == 0 {
		errs.FailWithJSON(c, "RoleIds和ID为必填")
		return
	}
	if err := h.service.Update(c, body); err != nil {
		errs.FailWithJSON(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}
