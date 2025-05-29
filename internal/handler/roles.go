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
// @Summary 创建角色
// @Tags 角色模块
// @Accept  json
// @Params data body model.Role
// @Success 200  {object} dto.Response[any]
// @Router /api/roles [post]
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

// List 角色列表
// @Summary 角色列表
// @Tags 角色模块
// @Accept  json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200  {object} dto.Response[dto.List[model.Role]]
// @Router /api/roles [get]
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
// @Summary 删除角色
// @Description 删除角色
// @Tags 角色模块
// @Accept  json
// @Param   id   path     int  true  "角色ID"
// @Param   data   body   dto.DeleteIds true "要删除的角色"
// @Success 200  {object} dto.Response[any]
// @Router /api/users/{id} [delete]
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
// @Summary 修改角色信息，支持name,description,和关联用户和关联权限表
// @Description 修改角色信息，支持name,description,和关联用户和关联权限表
// @Tags 角色模块
// @Accept  json
// @Param   id   path     int  true  "角色ID"
// @Param   data   body     dto.Role true "角色Id"
// @Success 200  {object} dto.Response[any]
// @Router /api/roles/{id} [put]
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

// UpdatePermissions 只修改角色权限
// @Summary 只修改角色权限
// @Description 只修改角色权限
// @Tags 角色模块
// @Accept  json
// @Param   id   path     int  true  "角色ID"
// @Param   data   body     dto.Role true "角色Id"
// @Success 200  {object} dto.Response[any]
// @Router /api/users/permissions/{id} [put]
func (h *RolesHandler) UpdatePermissions(c *gin.Context) {
	var roleId int
	id := c.Param("id")
	var role dto.Role
	if len(id) == 0 {
		errs.FailWithJSON(c, errs.New("id不能为空"))
		return
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if len(role.Permissions) == 0 {
		errs.FailWithJSON(c, errs.New("permissions不能为空"))
		return
	}

	if currentId, err := strconv.Atoi(id); err != nil {
		errs.FailWithJSON(c, err)
		return
	} else {
		roleId = currentId
	}

	if err := h.service.UpdatePermissions(c, roleId, role); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))

}
