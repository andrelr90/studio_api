package api

import (
	"net/http"
	"strconv"
	// "time"
	
	"github.com/gin-gonic/gin"

	"studio_api_project/main/models"
	"studio_api_project/main/repositories"
)


// Define API endpoints
func StartClassesAPI(router *gin.Engine) {
	router.GET("/classes", GetClasses)
	router.GET("/classes/:id", GetClass)
	router.POST("/classes", CreateClass)
	router.DELETE("/classes/:id", DeleteClass)
	router.PUT("/classes/:id", UpdateClass)
}

// GetClasses returns all classes
func GetClasses(c *gin.Context) {
	c.JSON(http.StatusOK, repositories.GetClasses())
}

// GetClass returns a specific class by ID
func GetClass(c *gin.Context) {
	id := c.Param("id")
	if class := repositories.GetClass(id); class == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
	} else {
		c.JSON(http.StatusOK, class)
	}
}

// CreateClass creates a new class
func CreateClass(c *gin.Context) {
	var class models.Class
	
	// // Convert the "start_date" and "end_date" fields to the RFC3339 format
	// rawData, err := c.GetRawData()
	// convertedData, err := models.ConvertDateFields(rawData, []string{"start_date", "end_date"})
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to convert date fields"})
	// 	return
	// }

	// // Bind the modified JSON data to the struct
	// err = json.Unmarshal(convertedData, &class)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to bind JSON data"})
	// 	return
	// }

	// Bind the modified JSON data to the struct
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate intersection with other classes
	if err := models.ValidateIntersection(repositories.GetClasses(), class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	class = *repositories.CreateClass(class)

	c.JSON(http.StatusCreated, class)
}

// DeleteClass deletes a class by ID
func DeleteClass(c *gin.Context) {
	id := c.Param("id")
	if err := repositories.DeleteClass(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
	}
}

// UpdateClass updates an existing class
func UpdateClass(c *gin.Context) {
	// Get the class ID from the request URL parameters
	classID := c.Param("id")
	
	// Find the class with the given ID
	class := repositories.GetClass(classID)
	if class == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	// Bind the modified JSON data to the struct
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Disallow changes in ID
	if (classID != strconv.Itoa(class.ID)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not allowed to change a class ID"})
		return
	}

	// Validate intersection with other classes
	if err := models.ValidateIntersection(repositories.GetClasses(), *class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the class in the storage or database (no need to check if the class exists, as it is tested by the first condition)
	updatedClass, _ := repositories.UpdateClassInStorage(class)

	c.JSON(http.StatusOK, updatedClass)
}
