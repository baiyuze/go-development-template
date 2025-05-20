package middleware

import "github.com/gin-gonic/gin"

type WhiteList struct {
	c *gin.Context
	// AuthWhiteList authWhiteList
}

func NewWhiteList(context *gin.Context) *WhiteList {
	return &WhiteList{
		c: context,
	}
}

// AuthWhiteList 认证白名单
func AuthWhiteList(c *gin.Context) {
	whiteList := []string{
		"public",
		"user/login",
		"user/register",
		//"user/logout",
	}
	c.Set("whiteList", whiteList)
	c.Next()
}
