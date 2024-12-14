package mysql

import (
	"fmt"
	"innoversepm-backend/pkg/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"innoversepm-backend/internal/setting"
)

var DB *gorm.DB

func InitMySQL(cfg *setting.AppConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.MySQL.Username,
		cfg.Database.MySQL.Password,
		cfg.Database.MySQL.Host,
		cfg.Database.MySQL.Port,
		cfg.Database.MySQL.Name,
	)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.BaseLogger.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// 获取底层的sql.DB对象
	sqlDB, err := DB.DB()
	if err != nil {
		logger.BaseLogger.Fatalf("Failed to get database instance: %v", err)
	}

	// 设置连接池参数，根据项目实际需求进行调整
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.BaseLogger.Println("MySQL connected successfully")
}

func CloseMySQL() {
	sqlDB, err := DB.DB()
	if err != nil {
		logger.BaseLogger.Fatalf("Failed to get database instance: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		logger.BaseLogger.Fatalf("Failed to close MySQL connection: %v", err)
	}
}
