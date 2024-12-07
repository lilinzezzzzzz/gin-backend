package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Config *AppConfig

// AppConfig 包含整个配置文件的结构
type AppConfig struct {
	App          App          `mapstructure:"app"`
	Database     Database     `mapstructure:"database"`
	Redis        Redis        `mapstructure:"redis"`
	DesignGPTAPI DesignGPTAPI `mapstructure:"design_gpt_api"`
}

// App 配置
type App struct {
	Port        int      `mapstructure:"port"`
	SecretKey   string   `mapstructure:"secret_key"`
	CORSOrigins []string `mapstructure:"cors_origins"`
}

// MySQL 配置
type MySQL struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

// Database 配置，包含多个数据库配置
type Database struct {
	MySQL          MySQL `mapstructure:"mysql"`
	DesignGPTMySQL MySQL `mapstructure:"design_gpt_mysql"`
}

// Redis 配置
type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// DesignGPTAPI 配置
type DesignGPTAPI struct {
	BaseURL string `mapstructure:"base_url"`
}

// LoadConfig 从 viper 加载配置到 Config 结构体
func LoadConfig(env string) {
	// 配置文件路径
	configFile := fmt.Sprintf("internal/config/%s.yaml", env)
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file %s: %v", configFile, err)
	}
	log.Printf("Using config file: %s", configFile)

	// 将配置映射到 Config 结构体
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	log.Println("Configuration loaded successfully")
}
