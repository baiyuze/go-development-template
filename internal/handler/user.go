package handler

import (
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
