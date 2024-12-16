package logger

import (
	"bytes"
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"innoversepm-backend/pkg/constants"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
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
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取时间戳并格式化
	timestamp := entry.Time.Format(time.RFC3339)

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
func InitLogrus(env string) {
	// 检查并创建 logs 目录
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create logs directory: %v", err)
		}
	}

	BaseLogger = logrus.New()
	var (
		formatter *CustomFormatter
		writer    io.Writer
	)
	switch env {
	case constants.DevEnvVal, constants.LocalEnvVal:
		// 创建带颜色的控制台 Formatter
		formatter = &CustomFormatter{EnableColor: true}
		writer = os.Stdout
	default:
		// 创建不带颜色的文件 Formatter
		formatter = &CustomFormatter{EnableColor: false}
		// 配置按天切割的日志文件
		logFile := &lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/app-%s.log", logDir, time.Now().Format("2006-01-02")),
			MaxSize:    10,   // 单个日志文件的最大大小（MB）
			MaxAge:     7,    // 保留旧日志的最大天数
			MaxBackups: 30,   // 保留旧日志的最大数量
			LocalTime:  true, // 使用本地时间
			Compress:   true, // 是否压缩旧日志
		}
		writer = logFile
	}

	// 设置 BaseLogger 同时输出到控制台和文件
	BaseLogger.SetOutput(io.MultiWriter(writer))
	// 启用调用者信息
	BaseLogger.SetReportCaller(true)
	BaseLogger.SetFormatter(formatter)
	BaseLogger.SetLevel(logrus.InfoLevel)
}
