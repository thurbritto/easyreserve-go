package database

import (
	"log"
	"os"

	"github.com/thurbritto/go-easyreserve/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// InitDB initializes the database connection and migrates the Reserve model.
func InitDB() {

	// Get the database connection string from environment variable
	dsn := os.Getenv("DATABASE_URL")

	// If the environment variable is not set, use a default connection string
	if dsn == "" {
		log.Fatal("Database connection string is not set")
		dsn = "root:123456@tcp(127.0.1:3306)/easyreserve?charset=utf8mb4&parseTime=True&loc=Local"
	}

	// Open a connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Check for connection errors
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Automigrate the Reserve model
	err = db.AutoMigrate(&models.Reserve{})

	// Check for migration errors
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Set the global DB variable
	DB = db
}
