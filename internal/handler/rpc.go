package handler

import (
	log "app/internal/common/log"
	"app/internal/grpc/client"
	"app/internal/grpc/container"
	"app/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type RpcHandler struct {
	service service.UserService
	clients *container.Clients
	log     *log.LoggerWithContext
}

func NewRpcHandler(
	service service.UserService,
	clients *container.Clients,
	log *log.LoggerWithContext,
) *RpcHandler {
	return &RpcHandler{
		service: service,
		clients: clients,
		log:     log,
	}
}

func ProviderRpcHandler(container *dig.Container) {
	err := container.Provide(NewRpcHandler)
	if err != nil {
		return
	}
}

// TestRpc 测试GRPC
// @Summary 测试GRPC
// @Tags GRPC
// @Accept  json
// @Router /api/rpc/test [get]
func (h *RpcHandler) TestRpc(c *gin.Context) {

	userValid, err := client.SayHello(h.clients)
	if err != nil {
		fmt.Println("查询失败:", err.Error())
	} else {
		fmt.Printf("查询数据: %+v\n", userValid)
	}
	c.JSON(http.StatusOK, userValid)
}
