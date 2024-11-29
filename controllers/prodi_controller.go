package controllers

import (
	"e-learning/database"
	"e-learning/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProdiInput struct {
	ID        string `json:"id" binding:"required"`
	NamaProdi string `json:"nama_prodi" binding:"required"`
}

type UpdateProdiInput struct {
	NamaProdi *string `json:"nama_prodi" binding:"required"`
}

func ShowProdi(c *gin.Context) {
	var prodi []models.Prodi

	// Fetch all Prodis from the database
	if err := database.DB.Find(&prodi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve prodi", "details": err.Error()})
		return
	}

	// Respond with the list of Prodis
	c.JSON(http.StatusOK, gin.H{
		"message": "Prodis retrieved successfully",
		"prodi":   prodi,
	})
}

func GetProdi(c *gin.Context) {
	prodiID := c.Param("id")
	id, err := strconv.Atoi(prodiID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Prodi ID"})
		return
	}

	var prodi models.Prodi
	if err := database.DB.Where("id = ?", id).Find(&prodi).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prodi not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Prodi retrieved successfully",
		"prodi":   prodi,
	})
}

func CreateProdi(c *gin.Context) {
	// Bind and validate input
	var input ProdiInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prodi := models.Prodi{
		IDProdi:   input.ID,
		NamaProdi: input.NamaProdi,
	}

	if err := database.DB.Create(&prodi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Prodi"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Prodi created successfully",
		"prodi":   prodi,
	})
}

func UpdateProdi(c *gin.Context) {
	// Get the Prodi ID from the URL
	prodiID := c.Param("id")

	var prodi models.Prodi
	if err := database.DB.Where("id = ?", prodiID).First(&prodi).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prodi not found"})
		return
	}

	// Bind the input JSON
	var input UpdateProdiInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.NamaProdi != nil {
		prodi.NamaProdi = *input.NamaProdi
	}

	// Save the updated Prodi
	if err := database.DB.Save(&prodi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Prodi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Prodi updated successfully", "Prodi": prodi})
}

func DeleteProdi(c *gin.Context) {
	// Get the Prodi ID from the URL
	prodiID := c.Param("id")

	var prodi models.Prodi
	if err := database.DB.Where("id = ?", prodiID).First(&prodi).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prodi not found"})
		return
	}

	// Delete the Prodi
	if err := database.DB.Delete(&prodi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Prodi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Prodi deleted successfully"})
}
