package handler

import (
	errs "app/internal/common/error"
	"app/internal/common/log"
	"app/internal/dto"
	"app/internal/grpc/container"
	"app/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"strconv"
)

type DepartmentHandler struct {
	service service.DepartmentService
	log     *log.LoggerWithContext
	clients *container.Clients
}

func NewDepartmentHandler(
	s service.DepartmentService,
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
	var body *dto.DepartmentBody
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
	var body *dto.DepartmentBody
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

// List tree查询
// @Summary 查询
// @Tags 部门
// @Accept  json
// @Success 200  {object} dto.Response[dto.List[model.Permission]]
// @Router /api/department [get]
func (h *DepartmentHandler) List(c *gin.Context) {
	tree, err := h.service.Tree(c)
	if err != nil {
		errs.FailWithJSON(c, err)
	} else {
		c.JSON(http.StatusOK, dto.Ok(tree))
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

// BindUser 绑定用户到部门
// @Summary 绑定用户到部门
// @Tags 部门
// @Accept  json
// @Params data params id
// @Params data body dto.UsersIds
// @Success 200  {object} dto.Response[any]
// @Router /api/department/{id}/users [post]
func (h *DepartmentHandler) BindUser(c *gin.Context) {
	var ids dto.UsersIds
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	if err := c.ShouldBindJSON(&ids); err != nil {
		errs.FailWithJSON(c, err)
	}

	if err := h.service.UpdateUsers(c, id, &ids); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}
