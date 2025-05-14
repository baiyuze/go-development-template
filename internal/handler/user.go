package handler

import (
	// AppContext "app/internal/app_ontext"

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
	container.Provide(NewUserHandler)
}

func (h *UserHandler) Login(c *gin.Context) {
	// logger := h.logx.WithContext(c)

	var body dto.LoginBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		errs.MustNoErr(err, "请检查账号密码")
	} else {

		result := h.service.Login(c, body)
		if result.Error != nil {
			errs.MustReturnErr(c, result.Error.Error())
			return
		}

		c.JSON(http.StatusOK, dto.Ok(gin.H{
			"token": result.Data,
		}))
	}
}

// HomeHandler 处理首页请求
func (h *UserHandler) HomeHandler(c *gin.Context) {
	// logger := h.logx.WithContext(c)

	user, err := h.service.GetUserOne()
	if err != nil {
		errs.MustNoErr(err, "错误了啊")
	} else {
		fmt.Printf("查询到的用户: %+v\n", user)
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) TestRpc(c *gin.Context) {

	userValid, err := client.SayHello(h.clients)
	if err != nil {
		fmt.Println("查询失败:", err.Error())
	} else {
		fmt.Printf("查询数据: %+v\n", userValid)
	}
	c.JSON(http.StatusOK, userValid)
}
