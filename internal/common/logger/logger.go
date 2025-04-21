package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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

// InitLogger 初始化生产环境日志，支持日志文件输出，并实现文件轮转
func InitLogger() (*zap.Logger, error) {
	var logFile = GetLogFilePath()
	// 使用 lumberjack 实现日志文件轮转
	writer := &lumberjack.Logger{
		Filename:   logFile, // 日志文件路径
		MaxSize:    10,      // MB, 每个日志文件最大大小
		MaxBackups: 3,       // 最多保留3个备份文件
		MaxAge:     28,      // 最长保存28天的日志
		Compress:   false,   // 是否压缩
	}
	localTime := time.Local
	// JSON 编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		// 使用本地时区（t.In(time.Local)）来格式化时间
		enc.AppendString(t.In(localTime).Format(time.RFC3339)) // 使用 RFC3339 格式并带本地时区
	}
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	// 创建文件同步器
	fileSyncer := zapcore.AddSync(writer)

	// 设置日志级别为 InfoLevel 或更高（生产环境通常不记录 Debug）
	level := zapcore.InfoLevel

	// 创建文件 core，文件使用 JSON 编码器
	fileCore := zapcore.NewCore(jsonEncoder, fileSyncer, level)

	// 创建 logger
	logger := zap.New(fileCore)

	return logger, nil
}

func GetLogFilePath() string {
	logName := "app"
	// 获取当前日期，用于日志文件名
	date := time.Now().Format("2006-01-02")

	// 获取用户的配置目录（跨平台）
	logDir, _ := os.UserConfigDir()

	// 日志文件路径：包含文件名
	logFilePath := filepath.Join(logDir, logName, "logs", fmt.Sprintf("%s-%s.log", logName, date))

	// 确保日志文件目录存在
	if _, err := os.Stat(filepath.Dir(logFilePath)); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(logFilePath), os.ModePerm)
		if err != nil {
			logFilePath = "/tmp/app.log" // 如果创建目录失败，退回到临时目录
		}
	}

	return logFilePath
}
