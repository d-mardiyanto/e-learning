package controllers

import (
	"e-learning/database"
	"e-learning/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassesInput struct {
	ID        string `json:"id" binding:"required"`
	ClassName string `json:"class_name" binding:"required"`
}

type UpdateClassesInput struct {
	ClassName *string `json:"class_name" binding:"required"`
}

func ShowClasses(c *gin.Context) {
	var classes []models.Classes

	// Fetch all Classes from the database
	if err := database.DB.Find(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Classes", "details": err.Error()})
		return
	}

	// Respond with the list of Classes
	c.JSON(http.StatusOK, gin.H{
		"message": "Classes retrieved successfully",
		"Classes": classes,
	})
}

func GetClasses(c *gin.Context) {
	ClassesID := c.Param("id")
	id, err := strconv.Atoi(ClassesID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Classes ID"})
		return
	}

	var classes models.Classes
	if err := database.DB.Where("id = ?", id).Find(&classes).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Classes not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Classes retrieved successfully",
		"Classes": classes,
	})
}

func CreateClasses(c *gin.Context) {
	// Bind and validate input
	var input ClassesInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	classes := models.Classes{
		ClassName: input.ClassName,
	}

	if err := database.DB.Create(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Classes"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Classes created successfully",
		"Classes": classes,
	})
}

func UpdateClasses(c *gin.Context) {
	// Get the Classes ID from the URL
	ClassesID := c.Param("id")

	var classes models.Classes
	if err := database.DB.Where("id = ?", ClassesID).First(&classes).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Classes not found"})
		return
	}

	// Bind the input JSON
	var input UpdateClassesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated Classes
	if err := database.DB.Save(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Classes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Classes updated successfully", "Classes": classes})
}

func DeleteClasses(c *gin.Context) {
	// Get the Classes ID from the URL
	ClassesID := c.Param("id")

	var classes models.Classes
	if err := database.DB.Where("id = ?", ClassesID).First(&classes).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Classes not found"})
		return
	}

	// Delete the Classes
	if err := database.DB.Delete(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Classes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Classes deleted successfully"})
}
