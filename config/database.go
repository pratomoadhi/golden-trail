package config

import (
	"fmt"
	"log"

	"github.com/pratomoadhi/golden-trail/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		panic("AutoMigrate User failed: " + err.Error())
	}

	if err := DB.AutoMigrate(&model.Transaction{}); err != nil {
		panic("AutoMigrate Transaction failed: " + err.Error())
	}
}
