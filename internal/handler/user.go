package handler

import (
	// AppContext "app/internal/app_ontext"

	errs "app/internal/common/error"
	"app/internal/common/logger"
	"app/internal/grpc/client"
	"app/internal/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// HomeHandler 处理首页请求
func (h *UserHandler) HomeHandler(c *gin.Context) {
	logger := logger.GetLogger(c)
	// l, exists := c.Get("logger")
	// if !exists {
	// 	c.JSON(500, gin.H{"error": "logger not found"})
	// 	return
	// }

	// logger := l.(*zap.Logger) // 类型断言
	logger.Info("测试")
	// var user models.User
	user, err := h.service.GetUserOne()
	errs.MustNoErr(errors.New("这是一个新错误"), "错误了啊")

	if err != nil {
		fmt.Println("查询失败:", err.Error())
	} else {
		fmt.Printf("查询到的用户: %+v\n", user)
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) TestRpc(c *gin.Context) {

	userValid, err := client.SayHello("嘻嘻")
	if err != nil {
		fmt.Println("查询失败:", err.Error())
	} else {
		fmt.Printf("查询数据: %+v\n", userValid)
	}
	c.JSON(http.StatusOK, userValid)
}
