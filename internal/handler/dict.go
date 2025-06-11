package handler

import (
	errs "app/internal/common/error"
	"app/internal/common/log"
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

type DictHandler struct {
	service service.DictService
	log     *log.LoggerWithContext
	clients *container.Clients
}

func NewDictHandler(
	s service.DictService,
	l *log.LoggerWithContext,
	clients *container.Clients,
) *DictHandler {
	return &DictHandler{
		service: s,
		log:     l,
		clients: clients,
	}
}

func ProviderDictHandler(container *dig.Container) {
	if err := container.Provide(NewDictHandler); err != nil {
		return
	}
}

// Create 创建
// @Summary 创建
// @Tags 字典管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Success 200  {object} dto.Response[any]
// @Router /api/dicts [post]
func (h *DictHandler) Create(c *gin.Context) {
	var body *model.Dict
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if len(body.Name) == 0 {
		errs.FailWithJSON(c, errs.New("name不能为空"))
		return
	}
	if len(body.Code) == 0 {
		errs.FailWithJSON(c, errs.New("code不能为空"))
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
// @Tags 字典管理
// @Accept  json
// @Param data body model.Dict true "body"
// @Success 200  {object} dto.Response[any]
// @Router /api/dicts/{id} [put]
func (h *DictHandler) Update(c *gin.Context) {
	var body *model.Dict
	var dictId int
	id := c.Param("id")
	if len(id) != 0 {
		result, err := strconv.Atoi(id)
		if err != nil {
			errs.FailWithJSON(c, err)
			return
		}
		dictId = result
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	if len(body.Name) == 0 {
		errs.FailWithJSON(c, errs.New("name不能为空"))
		return
	}
	// 不许更新code，更新code就得更新关联关系
	if len(body.Code) == 0 {
		errs.FailWithJSON(c, errs.New("code不能为空"))
		return
	}
	if err := h.service.Update(c, dictId, body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}

// List 查询
// @Summary 查询字典列表
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param name query string false "字典名称"
// @Success 200 {object} dto.Response[dto.List[[]model.Dict]]
// @Router /api/dicts [get]
func (h *DictHandler) List(c *gin.Context) {
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")
	name := c.Query("name")

	result, err := h.service.List(c, utils.HandleQuery(pageNum, pageSize), name)
	if err != nil {
		errs.FailWithJSON(c, err)
	} else {
		c.JSON(http.StatusOK, dto.Ok(result.Data))
	}
}

// GetOptions 根据Code获取options
// @Summary 根据Code获取options
// @Tags 字典管理
// @Accept  json
// @Param code query string true "字典编码"
// @Success 200  {object} dto.Response[[]model.DictItem]
// @Router /api/dicts/{code} [get]
func (h *DictHandler) GetOptions(c *gin.Context) {
	code := c.Param("code")
	dictItems, err := h.service.GetOptionsByDictCode(c, code)
	if err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Ok(dictItems))
}

// Delete 删除
// @Summary 删除
// @Tags 字典管理
// @Accept  json
// @Param data body dto.DeleteIds true "删除ids"
// @Success 200  {object} dto.Response[any]
// @Router /api/dicts [delete]
func (h *DictHandler) Delete(c *gin.Context) {
	var ids dto.DeleteIds
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
