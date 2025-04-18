package middleware

import (
	errs "app/internal/common/error"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryWithZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				traceId := "unknown"
				if v, ok := c.Get("traceId"); ok {
					if s, ok := v.(string); ok {
						traceId = s
					}
				}
				stack := string(debug.Stack())
				switch err := rec.(type) {
				case *errs.PanicError:
					logger.Error("业务 panic",
						zap.Int("code", err.Code),
						zap.String("message", err.Message),
						zap.Error(err.Err),
						zap.String("traceId", traceId),
						zap.Stack("stack"),
					)

					c.AbortWithStatusJSON(http.StatusOK, gin.H{
						"code":    err.Code,
						"msg":     err.Message,
						"traceId": traceId,
						"stack":   stack,
					})

				default:
					logger.Error("系统 panic",
						zap.Any("error", rec),
						zap.String("traceId", traceId),
						zap.Stack("stack"),
					)

					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"code":    500,
						"msg":     "internal server error",
						"traceId": traceId,
						"stack":   stack,
					})
				}
			}
		}()

		c.Next()
	}
}
