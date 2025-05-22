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

// Jwt 过滤白名单和验证token是否有效
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
	//白名单过滤
	if isHasPath {
		c.Next()
	} else {
		isPass := true
		var msg string
		if err := jwt.VerifyValidByToken(c, logger, "Authorization"); err != nil {
			msg = "Authorization verify token failed,err:" + err.Error()
			logger.Error("Authorization verify token failed", zap.Error(err))
			isPass = false
		}
		//如果token过期了，用refresh刷新token，refreshToken过期了，如果token没过期，刷新refreshToken
		if err := jwt.VerifyValidByToken(c, logger, "refreshToken"); err != nil {
			msg += ",refreshToken verify token failed,err:" + err.Error()
			logger.Error("refreshToken verify token failed", zap.Error(err))
			isPass = false
		}
		if isPass {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, dto.Fail(http.StatusUnauthorized, msg))
			c.Abort()
		}
	}
}
