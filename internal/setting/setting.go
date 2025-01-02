package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

var Config *AppConfig

// AppConfig 包含整个配置文件的结构
type AppConfig struct {
	App   App   `mapstructure:"app"`
	MySQL MySQL `mapstructure:"mysql"`
	Redis Redis `mapstructure:"redis"`
}

// App 配置
type App struct {
	Port        int      `mapstructure:"port"`
	SecretKey   string   `mapstructure:"secret_key"`
	CORSOrigins []string `mapstructure:"cors_origins"`
}

// MySQL 配置
type MySQL struct {
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	DBName          string `mapstructure:"dbname"`
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// Redis 配置
type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// LoadConfig 从 viper 加载配置到 Config 结构体
func LoadConfig(env string) {
	// 配置文件路径
	configFile := fmt.Sprintf("configs/%s.yaml", env)
	log.Printf("Using config file: %s", configFile)

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file %s: %v", configFile, err)
	}

	// 将配置映射到 Config 结构体
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	log.Println("Configuration loaded successfully")
}
