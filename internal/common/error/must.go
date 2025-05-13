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

func MustReturnErr(c *gin.Context, msg string) error {
	errJson := NewPanic(500, msg, nil)
	c.JSON(http.StatusForbidden, errJson)
	c.Abort() // 终止后续中间件和请求处理
	return nil

}
