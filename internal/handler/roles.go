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
	"strconv"
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

// Create 创建角色
func (h *RolesHandler) Create(c *gin.Context) {
	logger := h.log.WithContext(c)

	var body model.Role
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		logger.Error(err.Error())
		return
	}

	if len(body.Name) != 0 {

		if err := h.service.Create(c, body); err != nil {
			errs.FailWithJSON(c, err)
			return
		}
		c.JSON(http.StatusOK, dto.Ok[any](nil))
		return
	} else {
		errs.FailWithJSON(c, errs.New("角色名必填"))
		return
	}

}
func (h *RolesHandler) List(c *gin.Context) {
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")

	result, err := h.service.List(c, utils.HandleQuery(pageNum, pageSize))
	if err != nil {
		errs.FailWithJSON(c, err)
	} else {
		c.JSON(http.StatusOK, dto.Ok(result.Data))
	}
}

// Delete 删除角色
func (h *RolesHandler) Delete(c *gin.Context) {
	//logger := h.log.WithContext(c)
	var body dto.DeleteIds

	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if err := h.service.Delete(c, body); err != nil {

		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}

// Update 修改角色信息
func (h *RolesHandler) Update(c *gin.Context) {
	var roleId int
	id := c.Param("id")
	var role dto.Role
	if len(id) == 0 {
		errs.FailWithJSON(c, errs.New("id不能为空"))
		return
	}

	if currentId, err := strconv.Atoi(id); err != nil {
		errs.FailWithJSON(c, err)
		return
	} else {
		roleId = currentId
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	if len(role.Name) == 0 {
		errs.FailWithJSON(c, errs.New("name不能为空"))
		return
	}

	if err := h.service.Update(c, roleId, &role); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))

}

// UpdateRole 修改角色，设置角色
func (h *RolesHandler) UpdateRole(c *gin.Context) {
	var body dto.UserRoleRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	if len(body.RoleIds) == 0 || body.ID == 0 {
		errs.FailWithJSON(c, errs.New("RoleIds和ID为必填"))
		return
	}
	if err := h.service.UpdateRole(c, body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}
