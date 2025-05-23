package utils

import (
	"app/internal/dto"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	cp "github.com/otiai10/copy"
)

// 统一处理错误
func HandlerErr(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"message": err})
}

// 首字母大写
func CapitalizeFirstLetter(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// 首字母小写
func DecapitalizeFirstLetter(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false // 出错（例如路径不存在），默认返回 false
	}
	return info.IsDir() // 如果是目录，返回 true
}

// 读取文件后，并替换文件内容，同步
func ReplaceFileContent(sourceFile string, configArr []string, widgetName string) {
	name := configArr[0]
	entityName := configArr[1]
	oldName := configArr[2]
	oldEntityName := configArr[3]
	sFile, _ := filepath.Abs(sourceFile)
	// 插件名
	upName := CapitalizeFirstLetter(name)
	lowName := DecapitalizeFirstLetter(name)
	// 旧插件名
	upOldName := CapitalizeFirstLetter(oldName)
	lowOldName := DecapitalizeFirstLetter(oldName)
	// 新实体名
	upEntityName := CapitalizeFirstLetter(entityName)
	lowEntityName := DecapitalizeFirstLetter(entityName)
	// 旧实体名
	upOldEntityName := CapitalizeFirstLetter(oldEntityName)
	lowOldEntityName := DecapitalizeFirstLetter(oldEntityName)
	// tFile, _ := filepath.Abs(targetFile)
	content, err := os.ReadFile(sFile)
	if err != nil {
		return
	}

	newContentStr := strings.ReplaceAll(string(content), "${{widgetName}}", widgetName)
	newStr := strings.ReplaceAll(string(newContentStr), upOldName, upName)
	pluginStr := strings.ReplaceAll(newStr, lowOldName, lowName)
	entityStr := strings.ReplaceAll(pluginStr, upOldEntityName, upEntityName)
	entityLowStr := strings.ReplaceAll(entityStr, lowOldEntityName, lowEntityName)
	// 写入到文件中
	os.WriteFile(sourceFile, []byte(entityLowStr), 0777)
}

// 修改文件名称
func RenameFile(dir string, targetName string, oldName string, entity ...string) {
	slashDir := filepath.ToSlash(dir)
	targetDir := strings.ReplaceAll(slashDir, oldName, targetName)
	if len(entity) == 0 {

		if strings.Contains(targetDir, "MyPluginName") || strings.Contains(targetDir, "MyEntityName") {
			return
		}
		if !isDirectory(slashDir) {
			cp.Copy(slashDir, targetDir)
		}
	} else {
		targetEntityDir := strings.ReplaceAll(targetDir, entity[1], entity[0])
		if strings.Contains(targetEntityDir, "MyPluginName") || strings.Contains(targetEntityDir, "MyEntityName") {
			return
		}
		if !isDirectory(slashDir) {
			cp.Copy(slashDir, targetEntityDir)
		}
	}
}

// HandleQuery 处理参数query
func HandleQuery(pageNum string, pageSize string) dto.ListQuery {
	num, err := strconv.Atoi(pageNum)
	size, err := strconv.Atoi(pageSize)
	if pageNum == "" || err != nil {
		num = 1
	}
	if pageSize == "" || err != nil {
		size = 50
	}
	return dto.ListQuery{
		PageNum:  num,
		PageSize: size,
	}
}

// GetLocalIP 获取本地IP
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				if strings.Contains(ip4.String(), "192.168.") {
					return ip4.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("未找到非回环 IPv4 地址")
}
