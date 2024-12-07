package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Config *AppConfig

type AppConfig struct {
	App      App      `mapstructure:"app"`
	Database Database `mapstructure:"database"`
}

type App struct {
	Port int `mapstructure:"port"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// LoadConfig 从 viper 加载配置到 Config 结构体
func LoadConfig(env string) {
	// 相对路径
	configFile := fmt.Sprintf("internal/.env/%s.yaml", env)
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file %s: %v", configFile, err)
	}
	log.Printf("Using .env file: %s", configFile)

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	log.Println("Configuration loaded successfully")
}
