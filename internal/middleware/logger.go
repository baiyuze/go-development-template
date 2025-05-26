package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoggerMiddle struct {
	logger *zap.Logger
}

func NewLogger(logger *zap.Logger) *LoggerMiddle {
	return &LoggerMiddle{
		logger: logger,
	}
}

func (l *LoggerMiddle) Logger(c *gin.Context) {
	traceId, ok := c.Get(TraceIdKey)
	if !ok {
		l.logger.Error("TraceId不存在")
	} else {
		query := c.Request.URL.RawQuery
		loggerWithTrace := l.logger.With(
			zap.String("traceId", traceId.(string)),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", query))
		c.Set("logger", loggerWithTrace)
		loggerWithTrace.Info(fmt.Sprint(c.Request.URL.Path, " 🚀"),
			zap.String("ip", c.ClientIP()),
			zap.String("agent", c.Request.UserAgent()),
		)
		c.Next()
		
	}

}
