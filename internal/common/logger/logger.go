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

// InitLogger 初始化日志，控制台输出彩色日志，文件输出 JSON 并支持文件轮转
func InitLogger() (*zap.Logger, error) {
	logFile := GetLogFilePath()

	// 文件写入器（带轮转）
	fileWriter := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // MB
		MaxBackups: 3,
		MaxAge:     28, // 天
		Compress:   false,
	}
	fileSyncer := zapcore.AddSync(fileWriter)
	consoleSyncer := zapcore.AddSync(os.Stdout)

	// 设置日志级别
	level := zapcore.InfoLevel

	// 公共配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.In(time.Local).Format(time.RFC3339))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// JSON 编码器（用于文件）
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// Console 编码器（带颜色，用于终端）
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 构建两个 core：一个输出到文件，一个输出到控制台
	fileCore := zapcore.NewCore(jsonEncoder, fileSyncer, level)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleSyncer, level)

	// 合并 core
	core := zapcore.NewTee(
		fileCore,
		consoleCore,
	)

	// 构建 logger，带 caller 信息
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

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
