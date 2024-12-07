package setting

import (
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
func LoadConfig() {
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	log.Println("Configuration loaded successfully")
}
