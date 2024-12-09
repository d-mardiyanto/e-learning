package controllers

import (
	"e-learning/database"
	"e-learning/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudyProgramInput struct {
	ProgramName string `json:"program_name" binding:"required"`
}

func ShowStudyProgram(c *gin.Context) {
	var studyProgram []models.StudyProgram

	// Fetch all StudyPrograms from the database
	if err := database.DB.Find(&studyProgram).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve StudyProgram", "details": err.Error()})
		return
	}

	// Respond with the list of StudyPrograms
	c.JSON(http.StatusOK, gin.H{
		"message":      "StudyPrograms retrieved successfully",
		"StudyProgram": studyProgram,
	})
}

func GetStudyProgram(c *gin.Context) {
	StudyProgramID := c.Param("id")
	id, err := strconv.Atoi(StudyProgramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid StudyProgram ID"})
		return
	}

	var studyProgram models.StudyProgram
	if err := database.DB.Where("id = ?", id).Find(&studyProgram).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "StudyProgram not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "StudyProgram retrieved successfully",
		"StudyProgram": studyProgram,
	})
}

func CreateStudyProgram(c *gin.Context) {
	// Bind and validate input
	var input StudyProgramInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studyProgram := models.StudyProgram{
		ProgramName: input.ProgramName,
	}

	if err := database.DB.Create(&studyProgram).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create StudyProgram"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "StudyProgram created successfully",
		"StudyProgram": studyProgram,
	})
}

func UpdateStudyProgram(c *gin.Context) {
	// Get the StudyProgram ID from the URL
	StudyProgramID := c.Param("id")

	var studyProgram models.StudyProgram
	if err := database.DB.Where("id = ?", StudyProgramID).First(&studyProgram).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "StudyProgram not found"})
		return
	}

	// Bind the input JSON
	var input StudyProgramInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studyProgram.ProgramName = input.ProgramName

	// Save the updated StudyProgram
	if err := database.DB.Save(&studyProgram).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update StudyProgram"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "StudyProgram updated successfully", "StudyProgram": studyProgram})
}

func DeleteStudyProgram(c *gin.Context) {
	// Get the StudyProgram ID from the URL
	StudyProgramID := c.Param("id")

	var studyProgram models.StudyProgram
	if err := database.DB.Where("id = ?", StudyProgramID).First(&studyProgram).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "StudyProgram not found"})
		return
	}

	// Delete the StudyProgram
	if err := database.DB.Delete(&studyProgram).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete StudyProgram"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "StudyProgram deleted successfully"})
}
