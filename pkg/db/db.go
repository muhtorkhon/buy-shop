package db

import (
	"fmt"
	"i-shop/config"
	"i-shop/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Brand{}, &models.Category{}, &models.Product{}, &models.Image{}, &models.Order{}, &models.Users{}); err != nil {
		log.Fatal("Error Migratilon")
	}


	log.Println("Bazaga ulandi ")
	return db, nil
}

