package config

import (
	"log"
	"os"
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
	// viper.SetConfigFile(".env")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatalf("Error loading config: %v", err)
	// }

	// AppConfig = Config{
	// 	Port:       viper.GetString("PORT"),
	// 	DbHost:     viper.GetString("DB_HOST"),
	// 	DbPort:     viper.GetString("DB_PORT"),
	// 	DbUser:     viper.GetString("DB_USER"),
	// 	DbPassword: viper.GetString("DB_PASSWORD"),
	// 	DbName:     viper.GetString("DB_NAME"),
	// 	JwtSecret:  viper.GetString("JWT_SECRET"),
	// 	SentryDsn:  viper.GetString("SENTRY_DSN"),
	// }

	getEnv := func(key, defaultValue string) string {
		value := os.Getenv(key)
		if value == "" {
			if defaultValue == "" {
				log.Fatalf("Environment variable %s is required but not set", key)
			}
			return defaultValue
		}
		return value
	}

	AppConfig = Config{
		Port:       getEnv("PORT", "5000"),
		DbHost:     getEnv("DB_HOST", ""),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbUser:     getEnv("DB_USER", ""),
		DbPassword: getEnv("DB_PASSWORD", ""),
		DbName:     getEnv("DB_NAME", ""),
		JwtSecret:  getEnv("JWT_SECRET", ""),
		SentryDsn:  getEnv("SENTRY_DSN", ""),
	}

	return AppConfig
}
