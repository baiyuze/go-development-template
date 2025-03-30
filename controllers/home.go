package controllers

import (
	"app/dto"
	"app/service"
	"app/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// HomeHandler 处理首页请求
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

/**
* 读取文件
**/
func FileHandler(c *gin.Context) {
	query := c.Request.URL.Query()
	absPath, err := filepath.Abs("../" + query.Get("name"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	} else {
		path := filepath.ToSlash(absPath)
		envLocal, err := os.ReadFile(path)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": string(envLocal),
			})
		}
	}
}

// 创建组件
func CreateWidget(c *gin.Context) {
	var body dto.RequestBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		utils.HandlerErr(c, err)
		return
	}
	service.CreateService(c, body)

}
