package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thurbritto/go-easyreserve/database"
	"github.com/thurbritto/go-easyreserve/handlers"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	database.InitDB()

	// Initialize the Gin router
	router := gin.Default()

	// Set up a simple health check route
	router.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Register the reserve routes
	router.GET("/reserves", handlers.GetReserves)
	router.POST("/reserves", handlers.CreateReserve)

	// Start the server on port 8080
	router.Run(":8080")
}
