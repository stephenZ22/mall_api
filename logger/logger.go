package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	appLogger   *logrus.Logger
	errorLogger *logrus.Logger
)

// 初始化日志配置
func init() {
	// 创建并配置 appLogger
	appLogger = logrus.New()
	appLogger.SetFormatter(&logrus.TextFormatter{})
	appLogFile, err := os.OpenFile("logs/application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	appLogger.SetOutput(appLogFile)

	// 创建并配置 errorLogger
	errorLogger = logrus.New()
	errorLogger.SetFormatter(&logrus.JSONFormatter{})
	errorLogFile, err := os.OpenFile("logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		appLogger.Errorf("Failed to open error.log: %v", err)
	}
	errorLogger.SetOutput(errorLogFile)
}

// Error 输出错误日志到 error.log
func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

// Info 输出信息日志到 app.log
func Info(args ...interface{}) {
	appLogger.Info(args...)
}

// Warn 输出警告日志到 app.log
func Warn(args ...interface{}) {
	appLogger.Warn(args...)
}
