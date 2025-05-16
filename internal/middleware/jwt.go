package middleware

import (
	"app/internal/common/jwt"
	"app/internal/dto"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Jwt(c *gin.Context) {
	//先判断是否在白名单内
	log, ok := c.Get("logger")
	logger := log.(*zap.Logger)
	if !ok {
		fmt.Println("logger not found")
		return
	}
	whiteList, ok := c.Get("whiteList")
	var isHasPath bool
	if ok {
		for _, white := range whiteList.([]string) {
			if strings.Contains(c.Request.URL.Path, white) {
				isHasPath = true
				break
			} else {
				isHasPath = false
			}
		}
	}
	if isHasPath {
		c.Next()
	} else {
		tokenString := c.Request.Header.Get("Authorization")

		//// 先验证token有效性，再判断是否过期，如果过期，需要返回过期
		userInfo, err := jwt.Analysis(tokenString)

		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusUnauthorized, dto.Fail(http.StatusUnauthorized, err.Error()))
			c.Abort()
		} else {
			c.Set("userInfo", userInfo)
			c.Next()
		}

	}
}
