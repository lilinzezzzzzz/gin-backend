package config

import (
	"fmt"
	"log"
	"os"
)

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}
	configFile := fmt.Sprintf("internal/config/%s.yaml", env)
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file %s: %v", configFile, err)
	}
	log.Printf("Using config file: %s", configFile)

	Config = viper.GetViper()
}
