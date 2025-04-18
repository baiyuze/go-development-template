// internal/middleware/trace.go
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TraceIdKey = "traceId"

func Trace(c *gin.Context) {
	// 1. 生成 traceId（或者从 header 中读取）
	traceId := uuid.New().String()
	c.Set(TraceIdKey, traceId)

	// 2. 设置到响应头（前端也能看到）
	c.Writer.Header().Set("X-Trace-Id", traceId)

	c.Next()
}
