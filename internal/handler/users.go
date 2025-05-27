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

// SetRole 修改角色，设置角色
func (h *UserHandler) SetRole(c *gin.Context) {
	logger := h.log.WithContext(c)
	var body dto.UserRoleRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		logger.Error(err.Error())
		return
	}

	if len(body.RoleIds) == 0 || body.ID == 0 {
		errs.FailWithJSON(c, errs.New("RoleIds和ID为必填"))
		return
	}
	if err := h.service.Update(c, body); err != nil {
		errs.FailWithJSON(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.Ok[any](nil))
}

// TestAuth 用来验证是否token
func (h *UserHandler) TestAuth(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Ok("成功"))
}
