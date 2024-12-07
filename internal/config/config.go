package config

import "log"

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	Config = viper.GetViper()
}
