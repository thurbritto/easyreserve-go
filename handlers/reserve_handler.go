package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thurbritto/go-easyreserve/database"
	"github.com/thurbritto/go-easyreserve/models"
)

// GET /reserves
// GetReserves retrieves all reserves.
func GetReserves(c *gin.Context) {
	var reserves []models.Reserve
	if err := database.DB.Find(&reserves).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reserves."})
		return
	}

	// Return the list of reserves
	c.JSON(http.StatusOK, reserves)
}

// POST /reserves
// Createreserve creates a new reserve.
func CreateReserve(c *gin.Context) {

	// Define the input structure for reserve creation
	var input struct {
		UserID      int       `json:"user_id" binding:"required"`
		TableID     int       `json:"table_id" binding:"required"`
		ReserveDate time.Time `json:"reserve_date" binding:"required"`
		Guests      int       `json:"guests" binding:"required,min=1"`
		Status      string    `json:"status" binding:"required"`
	}

	// Validate the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate if the reserve date and time is in the future
	if input.ReserveDate.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reserve date and time must be in the future"})
		return
	}

	// Create a new reserve
	newReserve := models.Reserve{
		UserID:      input.UserID,
		TableID:     input.TableID,
		ReserveDate: input.ReserveDate,
		Guests:      input.Guests,
		Status:      input.Status,
		CreatedAt:   time.Now(),
	}

	// Try to save the reserve to the database
	if err := database.DB.Create(&newReserve).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save reserve."})
		return
	}

	// Return the created reserve
	c.JSON(http.StatusCreated, newReserve)
}
