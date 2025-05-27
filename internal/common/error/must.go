package errs

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AbortWithServerError 终止请求，系统错误，Panic
func AbortWithServerError(err error, msg string) {
	if err != nil {
		panic(NewPanic(500, msg, err))
	}
}

// FailWithJSON 请求失败，返回错误响应
func FailWithJSON(c *gin.Context, err error) {
	if l, exists := c.Get("logger"); exists {
		if logger, ok := l.(*zap.Logger); ok {
			logger.Error(err.Error(), zap.String("path", c.FullPath()))
		}
	}
	c.AbortWithStatusJSON(http.StatusForbidden, NewPanic(500, err.Error(), nil))
}
