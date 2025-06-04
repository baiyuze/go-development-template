package handler

import (
	errs "app/internal/common/error"
	"app/internal/common/log"
	"app/internal/dto"
	"app/internal/grpc/container"
	"app/internal/service"
	"app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"strconv"
)

type DepartmentHandler struct {
	service service.PermissionsService
	log     *log.LoggerWithContext
	clients *container.Clients
}

func NewDepartmentHandler(
	s service.PermissionsService,
	l *log.LoggerWithContext,
	clients *container.Clients,
) *DepartmentHandler {
	return &DepartmentHandler{
		service: s,
		log:     l,
		clients: clients,
	}
}

func ProviderDepartmentHandler(container *dig.Container) {
	if err := container.Provide(NewDepartmentHandler); err != nil {
		return
	}
}

// Create 创建
// @Summary 创建
// @Tags 部门
// @Accept  json
// @Params data body model.Permission
// @Success 200  {object} dto.Response[any]
// @Router /api/department [post]
func (h *DepartmentHandler) Create(c *gin.Context) {
	var body *dto.ReqPermissions
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if len(body.Name) == 0 {
		errs.FailWithJSON(c, errs.New("name不能为空"))
		return
	}
	if err := h.service.Create(c, body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}

// Update 更新
// @Summary 更新
// @Tags 部门
// @Accept  json
// @Params data body model.Permission
// @Success 200  {object} dto.Response[any]
// @Router /api/department [put]
func (h *DepartmentHandler) Update(c *gin.Context) {
	var body *dto.ReqPermissions
	var permissionId int
	id := c.Param("id")
	if len(id) != 0 {
		result, err := strconv.Atoi(id)
		if err != nil {
			errs.FailWithJSON(c, err)
			return
		}
		permissionId = result
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if len(body.Name) == 0 {
		errs.FailWithJSON(c, errs.New("name不能为空"))
		return
	}
	if err := h.service.Update(c, permissionId, body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}

// List 查询
// @Summary 查询
// @Tags 部门
// @Accept  json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200  {object} dto.Response[dto.List[model.Permission]]
// @Router /api/department [get]
func (h *DepartmentHandler) List(c *gin.Context) {
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")

	result, err := h.service.List(c, utils.HandleQuery(pageNum, pageSize))
	if err != nil {
		errs.FailWithJSON(c, err)
	} else {
		c.JSON(http.StatusOK, dto.Ok(result.Data))
	}
}

// Delete 删除
// @Summary 删除
// @Tags 部门
// @Accept  json
// @Params data body dto.DeleteIds
// @Success 200  {object} dto.Response[any]
// @Router /api/department [delete]
func (h *DepartmentHandler) Delete(c *gin.Context) {
	var ids *dto.DeleteIds
	if err := c.ShouldBindJSON(&ids); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if err := h.service.Delete(c, ids); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}
