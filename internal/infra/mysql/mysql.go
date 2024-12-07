package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"innoversepm-backend/internal/setting"
)

var DB *gorm.DB

func InitMySQL(cfg *setting.AppConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.MySQL.Username, cfg.Database.MySQL.Password, cfg.Database.MySQL.Host,
		cfg.Database.MySQL.Port, cfg.Database.MySQL.Name)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	log.Println("MySQL connected successfully")
}

// CloseMySQL 关闭数据库连接（如果需要）
func CloseMySQL() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close MySQL connection: %v", err)
	}
}
