package errs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MustNoErr(err error, msg string) {
	if err != nil {
		panic(NewPanic(500, msg, err))
	}
}

func MustReturnErr(c *gin.Context, err error, msg string) error {
	if err != nil {
		errJson := NewPanic(500, msg, err)
		c.JSON(http.StatusForbidden, errJson)
		c.Abort() // 终止后续中间件和请求处理
	}
	return nil

}
