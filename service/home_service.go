package service

import (
	"app/dto"
	"app/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	cp "github.com/otiai10/copy"
)

// CreateService 创建service
func CreateService(c *gin.Context, body dto.RequestBody) {
	// 复制文件到上一层
	absPath, _ := filepath.Abs("../src")
	tplAbsPath, _ := filepath.Abs("../src/MyPluginName")
	sourceAbsPath, _ := filepath.Abs("./template")
	menuAbsPath, _ := filepath.Abs("../src/config/menu.ts")
	widgetPath := filepath.ToSlash(absPath)
	sourcePath := filepath.ToSlash(sourceAbsPath)
	tplPath := filepath.ToSlash(tplAbsPath)
	upWidget := utils.CapitalizeFirstLetter(body.WidgetId)
	s, _ := filepath.Abs("../src/" + upWidget)
	t, _ := filepath.Abs("../src/widgets/" + upWidget)
	widgetSourcePath := filepath.ToSlash(s)
	widgetTargetPath := filepath.ToSlash(t)
	menuPath := filepath.ToSlash(menuAbsPath)

	cpErr := cp.Copy(sourcePath, widgetPath)
	if cpErr != nil {
		utils.HandlerErr(c, cpErr)
		return
	}

	// 扫描复制的文件夹
	var matchDirs []string
	err := filepath.WalkDir(tplPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			utils.HandlerErr(c, err)
			return err
		}
		matchDirs = append(matchDirs, path)
		return nil
	})
	if err != nil {
		utils.HandlerErr(c, err)
	}
	entityName := body.WidgetId + "Entity"
	for _, dir := range matchDirs {
		// 替换文件内容
		utils.ReplaceFileContent(dir, []string{body.WidgetId, entityName, "MyPluginName", "MyEntityName"}, body.WidgetName)
		// 修改文件名称
		if strings.Contains(dir, "MyPluginName") {
			if !strings.Contains(dir, "MyEntityName") {
				utils.RenameFile(dir, body.WidgetId, "MyPluginName")
			} else {
				utils.RenameFile(dir, body.WidgetId, "MyPluginName", entityName, "MyEntityName")
			}
		}

	}
	// 删除目录
	fileErr := os.RemoveAll(tplAbsPath)
	if fileErr != nil {
		utils.HandlerErr(c, fileErr)
	}
	// 移动目录，
	osErr := os.Rename(widgetSourcePath, widgetTargetPath)
	if osErr != nil {
		utils.HandlerErr(c, osErr)
	}
	// 更新 menu.ts
	newMenuMap := map[string]interface{}{
		"icon":      "p",
		"name":      body.WidgetName,
		"notPage":   false,
		"patchName": body.WidgetId,
		"path":      "/information-base/" + body.WidgetId,
	}
	body.Menu = append(body.Menu, newMenuMap)
	body.MenuMap[body.WidgetId] = newMenuMap
	menuStr, menuErr := generateMenuStr(body.Menu, 1)
	if menuErr != nil {
		utils.HandlerErr(c, menuErr)
	}
	menuMapStr, menuMapErr := generateMenuStr(body.MenuMap, 2)
	if menuMapErr != nil {
		utils.HandlerErr(c, menuMapErr)
	}
	menuSumStr := menuStr + menuMapStr
	// 写入文件
	writeErr := os.WriteFile(menuPath, []byte(menuSumStr), 0777)
	if writeErr != nil {
		utils.HandlerErr(c, writeErr)
	}
	appendFileContent(body.WidgetId, body.Type)
}

func generateMenuStr[T any](menuConfig T, menuType int) (string, error) {
	menuJson, menuErr := json.Marshal(menuConfig)
	var menuKey string
	var menuTypeKey string
	if menuType == 1 {
		menuKey = "menu"
		menuTypeKey = "Record<string,any>[]"
	} else {
		menuKey = "menuMap"
		menuTypeKey = "Record<string,any>"
	}
	str := fmt.Sprintf("export const %s: %s = ", menuKey, menuTypeKey)
	menuStr := str + string(menuJson) + "\n"
	return menuStr, menuErr
}

func appendFileContent(widgetId string, fileType int) (string, error) {
	absPath, _ := filepath.Abs("../.build.local")
	prodPath, _ := filepath.Abs("../.build.prod")
	local := filepath.ToSlash(absPath)
	prod := filepath.ToSlash(prodPath)

	var name string
	if fileType == 1 {
		name = local
	} else {
		name = prod
	}
	// 打开文件，如果文件不存在则创建
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {

		return widgetId, err
	}
	defer file.Close()

	// 追加内容到文件
	_, err = file.WriteString("\n" + widgetId + "\n")
	if err != nil {
		return widgetId, err
	}
	return widgetId, err
}
