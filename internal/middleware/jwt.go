package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Jwt(c *gin.Context) {
	//先判断是否在白名单内
	whiteList, ok := c.Get("white_List")
	fmt.Println(c.Request.URL.Path, whiteList, "--->")
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
		// 先验证token有效性，再判断是否过期，如果过期，需要返回过期

		c.Next()

	}
}
