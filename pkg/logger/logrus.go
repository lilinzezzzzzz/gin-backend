package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

var Logger *logrus.Logger

// CustomFormatter 自定义日志格式
type CustomFormatter struct{}

// Format 实现 logrus.Formatter 接口
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取时间戳并格式化
	timestamp := entry.Time.Format("2006-01-02 15:04:05 -0700")

	// 获取日志级别
	level := entry.Level.String()

	// 获取调用者文件和行号（需要启用 ReportCaller）
	var file string
	if entry.Caller != nil {
		file = fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	} else {
		file = "unknown:0"
	}

	// 获取 trace_id，如果没有则设置为空字符串
	traceID, ok := entry.Data["trace_id"]
	if !ok {
		traceID = ""
	}

	// 日志消息
	message := entry.Message

	// 拼接日志内容
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%s|%s|%s|%v|%s\n", timestamp, level, file, traceID, message))

	return b.Bytes(), nil
}

// InitLogrus 初始化Logrus日志配置
func InitLogrus() {
	Logger = logrus.New()

	// 启用调用者信息
	Logger.SetReportCaller(true)

	// 设置自定义 Formatter
	Logger.SetFormatter(&CustomFormatter{})

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
