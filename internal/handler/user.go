package handler

import (
	// AppContext "app/internal/app_ontext"

	errs "app/internal/common/error"
	"app/internal/common/logx"
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
	logx    *logx.LoggerWithContext
}

func NewUserHandler(
	service service.UserService,
	clients *container.Clients,
	logx *logx.LoggerWithContext,
) *UserHandler {
	return &UserHandler{
		service: service,
		clients: clients,
		logx:    logx,
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

		sign, err := h.service.Login(c, body)
		if err != nil {
			errs.MustReturnErr(c, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"token": sign,
			},
			"msg": nil,
		})
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
