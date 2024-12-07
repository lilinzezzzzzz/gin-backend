package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger *logrus.Logger

// InitLogrus 初始化Logrus日志配置
func InitLogrus() {
	Logger := logrus.New()

	// 设置日志格式为JSON格式
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志输出到控制台和文件
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		Logger.SetOutput(mw)
	} else {
		Logger.Info("Failed to log to file, using default stderr")
	}

	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)
}
