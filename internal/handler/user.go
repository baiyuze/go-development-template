package handler

import (
	"app/internal/grpc/client"
	"app/internal/service"
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
	// var user models.User
	user, err := h.service.GetUserOne()
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println("查询失败:", err.Error())
	} else {
		fmt.Printf("查询到的用户: %+v\n", user)
	}
	c.JSON(http.StatusOK, user)
}

var userClient = client.NewHelloClient()

func (h *UserHandler) TestRpc(c *gin.Context) {
	// var user models.User
	fmt.Printf("-----<")
	userValid, err := userClient.SayHello("test")

	if err != nil {
		fmt.Println("查询失败:", err.Error())
	} else {
		fmt.Printf("查询数据: %+v\n", userValid)
	}
	c.JSON(http.StatusOK, userValid)
}
