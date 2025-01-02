package infra

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang-backend/internal/setting"
	pkgLogger "golang-backend/pkg/logger"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB 全局数据库实例
var DB *gorm.DB

func NewDB(ctx *gin.Context) *gorm.DB {
	traceID := ctx.GetString("trace_id")
	// 先把 gin.Context 转为标准 context，或者直接 ctx.Request.Context()
	stdCtx := ctx.Request.Context()
	// 3. 手动把 trace_id 塞进去
	stdCtx = context.WithValue(stdCtx, "trace_id", traceID)
	// 在同一个链式调用里加上 context 和 trace_id
	db := DB.WithContext(stdCtx)
	return db
}

// InitMySQL 初始化 MySQL 数据库连接
func InitMySQL(cfg *setting.AppConfig) {
	// 构建 DSN (Data Source DBName)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.Username,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DBName,
	)

	// 配置 GORM 日志级别
	//newLogger := logger.New(
	//	logger.New(logger.Writer(), "\r\n", logger.LstdFlags), // 使用标准日志库的输出
	//	logger.Config{
	//		SlowThreshold:             time.Second, // 慢查询阈值
	//		LogLevel:                  logger.Info, // 日志级别
	//		IgnoreRecordNotFoundError: true,        // 忽略未找到记录的错误
	//		Colorful:                  true,        // 彩色打印
	//	},
	//)
	newLogger := &GormLogger{
		logger: pkgLogger.BaseLogger,
		cfg: logger.Config{
			SlowThreshold:             time.Second * 2, // 慢查询阈值
			LogLevel:                  logger.Info,     // GORM 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略未找到记录的错误
			Colorful:                  true,            // 彩色打印
		},
	}

	// 连接数据库
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加复数
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// 获取底层的 *sql.DB，用于设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(cfg.MySQL.MaxOpenConns)       // 最大打开的连接数
	sqlDB.SetMaxIdleConns(cfg.MySQL.MaxIdleConns)       // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(cfg.MySQL.ConnMaxLifetime) // 连接最大生命周期

	log.Println("MySQL database connected successfully")
}

// CloseMySQL 关闭数据库连接
func CloseMySQL() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("Failed to get database instance: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
		log.Println("MySQL database connection closed")
	}
}

// GormLogger 这是自定义的 GORM Logger
type GormLogger struct {
	logger *logrus.Logger
	cfg    logger.Config
}

// 确保我们实现了 gorm.io/gorm/logger.Interface
var _ logger.Interface = (*GormLogger)(nil)

// LogMode 改变日志级别时被调用，可以按需实现
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.cfg.LogLevel = level
	return &newLogger
}

// Info 这三个可以根据需要实现或留空
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{})  {}
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{})  {}
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {}

// Trace 关键方法：GORM 每执行一次 SQL，都会调用这里
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	// 1) 取出 trace_id（若你在 handler 里用 context.WithValue("trace_id", ...)）
	traceID, _ := ctx.Value("trace_id").(string)

	// 2) 处理错误、慢查询等逻辑
	switch {
	case err != nil:
		// 打印带有 error, trace_id
		l.logger.WithFields(logrus.Fields{
			"trace_id": traceID,
			"elapsed":  elapsed,
			"rows":     rows,
		}).WithError(err).Errorf("SQL => %s", sql)
	// 如果你想要慢查询日志
	case elapsed > l.cfg.SlowThreshold && l.cfg.SlowThreshold != 0:
		l.logger.WithFields(logrus.Fields{
			"trace_id": traceID,
			"elapsed":  elapsed,
			"rows":     rows,
		}).Warnf("SLOW SQL => %s", sql)
	default:
		// 正常 SQL 日志
		l.logger.WithFields(logrus.Fields{
			"trace_id": traceID,
			"elapsed":  elapsed,
			"rows":     rows,
		}).Infof("SQL => %s", sql)
	}
}
