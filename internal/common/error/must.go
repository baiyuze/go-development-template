package errs

import (
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
func FailWithJSON(c *gin.Context, msg string) {
	errJson := NewPanic(500, msg, nil)
	c.JSON(http.StatusForbidden, errJson)
	c.Abort() // 终止后续中间件和请求处理
}
