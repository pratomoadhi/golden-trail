package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	JwtSecret  string
	SentryDsn  string
}

var AppConfig Config

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	AppConfig = Config{
		Port:       viper.GetString("PORT"),
		DbHost:     viper.GetString("DB_HOST"),
		DbPort:     viper.GetString("DB_PORT"),
		DbUser:     viper.GetString("DB_USER"),
		DbPassword: viper.GetString("DB_PASSWORD"),
		DbName:     viper.GetString("DB_NAME"),
		JwtSecret:  viper.GetString("JWT_SECRET"),
		SentryDsn:  viper.GetString("SENTRY_DSN"),
	}

	return AppConfig
}
