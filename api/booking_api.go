package api

import (
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"

	"studio_api_project/main/models"
	"studio_api_project/main/repositories"
)

// Define API endpoints
func StartBookingsAPI(router *gin.Engine) {
	router.GET("/bookings", GetBookings)
	router.GET("/bookings/:id", GetBooking)
	router.POST("/bookings", CreateBooking)
	router.DELETE("/bookings/:id", DeleteBooking)
	router.PUT("/bookings/:id", UpdateBooking)
}

func GetBookings(c *gin.Context) {
	c.JSON(http.StatusOK, repositories.GetBookings())
}

func GetBooking(c *gin.Context) {
	id := c.Param("id")
	if booking := repositories.GetBooking(id); booking == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
	} else {
		c.JSON(http.StatusOK, booking)
	}
}

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Validate creation:
	createdBooking := repositories.CreateBooking(booking)
	if createdBooking == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There are no classes in this date"})
		return
	}

	c.JSON(http.StatusCreated, createdBooking)
}

func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	if err := repositories.DeleteBooking(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Class booking deleted successfully"})
	}
}

func UpdateBooking(c *gin.Context) {
	// Get the class booking ID from the request URL parameters
	bookingID := c.Param("id")

	// Find the class booking with the given ID
	booking := repositories.GetBooking(bookingID)
	if booking == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Bind the request JSON data to the class booking object
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Disallow changes in ID
	if (bookingID != strconv.Itoa(booking.ID)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not allowed to change a class booking ID"})
		return
	}

	// Update the class in the storage or database
	updatedBooking, err := repositories.UpdateBookingInStorage(booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBooking)
}
