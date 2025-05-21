package handler

import (
	errs "app/internal/common/error"
	log "app/internal/common/log"
	"app/internal/dto"
	"app/internal/grpc/client"
	"app/internal/grpc/container"
	"app/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
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

func (h *UserHandler) Login(c *gin.Context) {
	logger := h.log.WithContext(c)

	var body dto.LoginBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		errs.AbortWithServerError(err, "请检查账号密码")
	} else {

		result := h.service.Login(c, body)
		if result.Error != nil {
			errs.FailWithJSON(c, result.Error.Error())
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
		errs.FailWithJSON(c, err.Error())
		logger.Error(err.Error())
		return
	}

	account := body.Account
	if account != nil || body.Password != nil {
		if err := h.service.Register(c, body); err != nil {
			errs.FailWithJSON(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, dto.Ok[any](nil))
		return
	} else {
		errs.FailWithJSON(c, "账号或密码不存在")
		return
	}

	//fmt.Printf("%+v", body, "--->")

}

// TestAuth 用来验证是否token
func (h *UserHandler) TestAuth(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Ok("成功"))
}

// HomeHandler 处理首页请求
func (h *UserHandler) HomeHandler(c *gin.Context) {
	// logger := h.logx.WithContext(c)

	user, err := h.service.GetUserOne()
	if err != nil {
		errs.AbortWithServerError(err, "错误了啊")
	} else {
		fmt.Printf("查询到的用户: %+v\n", user)
	}
	c.JSON(http.StatusOK, user)
}

// TestRpc 测试GRPC
func (h *UserHandler) TestRpc(c *gin.Context) {

	userValid, err := client.SayHello(h.clients)
	if err != nil {
		fmt.Println("查询失败:", err.Error())
	} else {
		fmt.Printf("查询数据: %+v\n", userValid)
	}
	c.JSON(http.StatusOK, userValid)
}
