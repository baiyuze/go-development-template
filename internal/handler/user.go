package handler

import (
	// AppContext "app/internal/app_ontext"

	errs "app/internal/common/error"
	"app/internal/common/logger"
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
}

func NewUserHandler(service service.UserService, clients *container.Clients) *UserHandler {
	return &UserHandler{
		service: service,
		clients: clients,
	}
}

func ProviderUserHandler(container *dig.Container) {
	container.Provide(NewUserHandler)
}

// HomeHandler 处理首页请求
func (h *UserHandler) HomeHandler(c *gin.Context) {
	logger := logger.GetLogger(c)

	logger.Info("测试")
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
