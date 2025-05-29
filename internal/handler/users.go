package handler

import (
	errs "app/internal/common/error"
	log "app/internal/common/log"
	"app/internal/dto"
	"app/internal/grpc/container"
	"app/internal/service"
	"app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service service.UserService
	clients *container.Clients
	log     *log.LoggerWithContext
}

func NewUserHandler(
	service service.UserService,
	clients *container.Clients,
	log *log.LoggerWithContext,
) *UserHandler {
	return &UserHandler{
		service: service,
		clients: clients,
		log:     log,
	}
}

func ProviderUserHandler(container *dig.Container) {
	err := container.Provide(NewUserHandler)
	if err != nil {
		return
	}
}

// HomeHandler 处理首页请求
func (h *UserHandler) HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "首页")
}

// Login 登录接口
// @Summary 登录接口
// @Tags 用户模块
// @Accept  json
// @Param   data  body dto.LoginBody  true  "用户信息"
// @Success 200  {object} dto.Response[any]
// @Router /api/users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	logger := h.log.WithContext(c)

	var body dto.LoginBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		errs.AbortWithServerError(err, "请检查账号密码")
	} else {

		result := h.service.Login(c, body)
		if result.Error != nil {
			errs.FailWithJSON(c, result.Error)
			logger.Error(result.Error.Error())
			return
		}

		c.JSON(http.StatusOK, dto.Ok(gin.H{
			"token": result.Data,
		}))
	}
}

// Register 注册
// @Summary 注册
// @Tags 用户模块
// @Accept  json
// @Param   data  body dto.RegBody  true  "注册用户"
// @Success 200  {object} dto.Response[any]
// @Router /api/users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	logger := h.log.WithContext(c)

	var body dto.RegBody
	if err := c.ShouldBindJSON(&body); err != nil {
		errs.FailWithJSON(c, err)
		logger.Error(err.Error())
		return
	}

	account := body.Account
	if account != nil || body.Password != nil {
		if err := h.service.Register(c, body); err != nil {
			errs.FailWithJSON(c, err)
			return
		}
		c.JSON(http.StatusOK, dto.Ok[any](nil))
		return
	} else {
		errs.FailWithJSON(c, errs.New("账号或密码不存在"))
		return
	}

}

// List 用户列表
// @Summary 用户列表
// @Tags 用户模块
// @Accept  json
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200  {object} dto.Response[dto.List[dto.UserWithRole]]
// @Router /api/users [get]
func (h *UserHandler) List(c *gin.Context) {
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")

	result, err := h.service.List(c, utils.HandleQuery(pageNum, pageSize))
	if err != nil {
		errs.FailWithJSON(c, err)
	} else {
		c.JSON(http.StatusOK, dto.Ok(result.Data))
	}
}

// UpdateRole 修改角色，设置角色
// @Summary 设置角色
// @Description 修改角色，设置角色
// @Tags 用户模块
// @Accept  json
// @Param   id   path     int  true  "用户ID"
// @Success 200  {object} dto.User
// @Router /api/users/{id} [put]
func (h *UserHandler) UpdateRole(c *gin.Context) {
	var userId int
	id := c.Param("id")
	var user dto.User
	if len(id) == 0 {
		errs.FailWithJSON(c, errs.New("id不能为空"))
		return
	}

	if currentId, err := strconv.Atoi(id); err != nil {
		errs.FailWithJSON(c, err)
		return
	} else {
		userId = currentId
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		errs.FailWithJSON(c, err)
		return
	}

	if err := h.service.UpdateRoles(c, userId, &user); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}

// Delete 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags 用户模块
// @Accept  json
// @Param   id   path     int  true  "用户ID"
// @Success 200  {object} dto.DeleteIds
// @Router /api/users [delete]
func (h *UserHandler) Delete(c *gin.Context) {
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

// TestAuth 用来验证是否token
func (h *UserHandler) TestAuth(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Ok("成功"))
}
