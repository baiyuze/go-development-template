package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetLogger 获取logger
func GetLogger(c *gin.Context) *zap.Logger {
	l, exists := c.Get("logger")
	if !exists {
		return zap.NewNop() // 返回一个不会打印的空 logger，避免 panic
	}
	if logger, ok := l.(*zap.Logger); ok {
		return logger
	}
	return zap.NewNop()
}
