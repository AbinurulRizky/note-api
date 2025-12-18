package config

import (
	"fmt"
	"os"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"note/app/models"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Attempt %d: Failed to connect to database: %v", i+1, err)
			time.Sleep(2 * time.Second)
		} else {
			log.Println("Successfully connected to the database")
		}
	}

	if err != nil {
		log.Fatalf("Could not connect to the database after multiple attempts: %v", err)
	}

	// Auto-migrate models
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	return DB
}