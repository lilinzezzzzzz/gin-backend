package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

var BaseLogger *logrus.Logger

// 定义颜色代码
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// CustomFormatter 自定义日志格式
type CustomFormatter struct {
	EnableColor bool // 是否启用颜色
}

// Format 实现 logrus.Formatter 接口
// Format 实现 logrus.Formatter 接口
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取时间戳并格式化
	timestamp := entry.Time.Format("2006-01-02 15:04:05 -0700")

	// 获取日志级别和对应颜色
	level := strings.ToUpper(entry.Level.String())
	var levelColor string

	if f.EnableColor {
		switch entry.Level {
		case logrus.DebugLevel:
			levelColor = Cyan
		case logrus.InfoLevel:
			levelColor = Green
		case logrus.WarnLevel:
			levelColor = Yellow
		case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
			levelColor = Red
		default:
			levelColor = White
		}
	}

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
	if f.EnableColor {
		b.WriteString(fmt.Sprintf("|%s%s%s|%s|%s|%s|%s\n", levelColor, level, Reset, timestamp, traceID, file, message))
	} else {
		b.WriteString(fmt.Sprintf("|%s|%s|%s|%s|%s\n", level, timestamp, traceID, file, message))
	}

	return b.Bytes(), nil
}

// InitLogrus 初始化 Logrus 日志配置
// InitLogrus 初始化 Logrus 日志配置
func InitLogrus() {
	BaseLogger = logrus.New()

	// 启用调用者信息
	BaseLogger.SetReportCaller(true)

	// 创建带颜色的控制台 Formatter
	consoleFormatter := &CustomFormatter{EnableColor: true}

	// 创建不带颜色的文件 Formatter
	fileFormatter := &CustomFormatter{EnableColor: false}

	// 设置控制台输出
	consoleOutput := os.Stdout
	consoleLogger := logrus.New()
	consoleLogger.SetOutput(consoleOutput)
	consoleLogger.SetFormatter(consoleFormatter)
	consoleLogger.SetLevel(logrus.InfoLevel)
	consoleLogger.SetReportCaller(true)

	// 设置文件输出
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("set file err: %v", err.Error())
	}

	fileLogger := logrus.New()
	fileLogger.SetOutput(file)
	fileLogger.SetFormatter(fileFormatter)
	fileLogger.SetLevel(logrus.InfoLevel)
	fileLogger.SetReportCaller(true)

	// 多路输出到控制台和文件
	BaseLogger.SetOutput(io.MultiWriter(consoleOutput, file))
	BaseLogger.SetFormatter(consoleFormatter)
}
