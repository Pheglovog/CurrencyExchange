package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Host string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var AppConfig *Config

func InitConfig() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	// If a configs file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}

	// Unmarshal the config into the struct
	AppConfig = &Config{}
	err = viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	InitDB()
	InitRedis()
}
