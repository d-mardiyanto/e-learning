package controllers

import (
	"e-learning/database"
	"e-learning/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InstructorsInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UpdateInstructorsInput struct {
	Name  *string `json:"name" binding:"required"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

func ShowInstructors(c *gin.Context) {
	var instructors []models.Instructors

	// Fetch all Instructors from the database
	if err := database.DB.Find(&instructors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Instructors", "details": err.Error()})
		return
	}

	// Respond with the list of Instructorss
	c.JSON(http.StatusOK, gin.H{
		"message":     "Instructors retrieved successfully",
		"instructors": instructors,
	})
}

func GetInstructor(c *gin.Context) {
	InstructorsID := c.Param("id")
	id, err := strconv.Atoi(InstructorsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Instructors ID"})
		return
	}

	var instructors models.Instructors
	if err := database.DB.Where("id = ?", id).Find(&instructors).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructors not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Instructors retrieved successfully",
		"instructors": instructors,
	})
}

func CreateInstructor(c *gin.Context) {
	// Bind and validate input
	var input InstructorsInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructors := models.Instructors{
		Nama:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	if err := database.DB.Create(&instructors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Instructors"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Instructors created successfully",
		"instructors": instructors,
	})
}

func UpdateInstructor(c *gin.Context) {
	// Get the Instructors ID from the URL
	InstructorsID := c.Param("id")

	var instructors models.Instructors
	if err := database.DB.Where("id = ?", InstructorsID).First(&instructors).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructors not found"})
		return
	}

	// Bind the input JSON
	var input UpdateInstructorsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != nil {
		instructors.Nama = *input.Name
	}
	if input.Email != nil {
		instructors.Email = *input.Email
	}
	if input.Phone != nil {
		instructors.Phone = *input.Phone
	}

	// Save the updated Instructors
	if err := database.DB.Save(&instructors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Instructors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Instructors updated successfully", "instructors": instructors})
}

func DeleteInstructor(c *gin.Context) {
	// Get the Instructors ID from the URL
	InstructorsID := c.Param("id")

	var instructors models.Instructors
	if err := database.DB.Where("id = ?", InstructorsID).First(&instructors).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructors not found"})
		return
	}

	// Delete the Instructors
	if err := database.DB.Delete(&instructors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Instructors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Instructors deleted successfully"})
}
