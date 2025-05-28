package middleware

import (
	"app/internal/common/jwt"
	"app/internal/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
)

// Jwt 过滤白名单和验证token是否有效
func Jwt(verifyToken bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		//dev开发环境下，可以不传token
		env := os.Getenv("ENV")
		if env != "production" {
			c.Next()
			return
		}
		//去除白名单方式校验token，不够优雅
		log, ok := c.Get("logger")
		logger := log.(*zap.Logger)
		if !ok {
			fmt.Println("logger not found")
			return
		}

		if !verifyToken {
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

}
