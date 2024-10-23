package config

import (
	"log"
	"TaskHandler/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	// Membuka koneksi SQLite
	database, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Melakukan migrasi model
	if err := database.AutoMigrate(&models.Book{}); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	return database
}
