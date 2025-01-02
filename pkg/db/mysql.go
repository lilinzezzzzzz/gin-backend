package db

import (
	"fmt"
	pkgLogger "golang-backend/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// DB 全局数据库实例
var DB *gorm.DB

// InitMySQL 初始化 MySQL 数据库连接
func InitMySQL(username, password, host, port, dbName string, maxOpenConn, maxIdleConn int, connMaxLifetime time.Duration) {
	// 构建 DSN (Data Source DBName)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName,
	)

	// 配置 GORM 日志级别
	newLogger := &pkgLogger.SQLLogger{
		Logger: pkgLogger.BaseLogger,
		Cfg: gormLogger.Config{
			SlowThreshold:             time.Second * 2, // 慢查询阈值
			LogLevel:                  gormLogger.Info, // GORM 日志级别
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

	// 获取底层的 *sql.db，用于设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(maxOpenConn)        // 最大打开的连接数
	sqlDB.SetMaxIdleConns(maxIdleConn)        // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(connMaxLifetime) // 连接最大生命周期

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
