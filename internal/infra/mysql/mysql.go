package mysql

import (
	"fmt"
	"innoversepm-backend/internal/setting"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB 全局数据库实例
var DB *gorm.DB

// InitMySQL 初始化 MySQL 数据库连接
func InitMySQL(cfg *setting.AppConfig) {
	// 构建 DSN (Data Source DBName)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.MySQL.Username,
		cfg.Database.MySQL.Password,
		cfg.Database.MySQL.Host,
		cfg.Database.MySQL.Port,
		cfg.Database.MySQL.DBName)

	// 配置 GORM 日志级别
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // 使用标准日志库的输出
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略未找到记录的错误
			Colorful:                  true,        // 彩色打印
		},
	)

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
	sqlDB.SetMaxOpenConns(cfg.Database.MySQL.MaxOpenConns)       // 最大打开的连接数
	sqlDB.SetMaxIdleConns(cfg.Database.MySQL.MaxIdleConns)       // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(cfg.Database.MySQL.ConnMaxLifetime) // 连接最大生命周期

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
