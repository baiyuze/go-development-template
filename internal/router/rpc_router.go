package router

import (
	"app/internal/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterRpcRoutes 注册所有路由
func RegisterRpcRoutes(r *gin.Engine, container *dig.Container) {

	router := r.Group("rpc")
	err := container.Invoke(func(rpcHandler *handler.RpcHandler) {
		// 测试RPC
		router.GET("/test", rpcHandler.TestRpc)
	})
	if err != nil {
		fmt.Printf("注入 handler 失败: %v\n", err)
	}

}
