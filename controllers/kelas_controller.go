package controllers

import (
	"e-learning/database"
	"e-learning/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KelasInput struct {
	ID        string `json:"id" binding:"required"`
	NamaKelas string `json:"nama_kelas" binding:"required"`
}

type UpdateKelasInput struct {
	NamaKelas *string `json:"nama_kelas" binding:"required"`
}

func ShowKelas(c *gin.Context) {
	var kelas []models.Kelas

	// Fetch all Kelass from the database
	if err := database.DB.Find(&kelas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Kelas", "details": err.Error()})
		return
	}

	// Respond with the list of Kelass
	c.JSON(http.StatusOK, gin.H{
		"message": "Kelass retrieved successfully",
		"kelas":   kelas,
	})
}

func GetKelas(c *gin.Context) {
	KelasID := c.Param("id")
	id, err := strconv.Atoi(KelasID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Kelas ID"})
		return
	}

	var kelas models.Kelas
	if err := database.DB.Where("id = ?", id).Find(&kelas).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kelas not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kelas retrieved successfully",
		"kelas":   kelas,
	})
}

func CreateKelas(c *gin.Context) {
	// Bind and validate input
	var input KelasInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kelas := models.Kelas{
		IDKelas:   input.ID,
		NamaKelas: input.NamaKelas,
	}

	if err := database.DB.Create(&kelas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Kelas"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Kelas created successfully",
		"kelas":   kelas,
	})
}

func UpdateKelas(c *gin.Context) {
	// Get the Kelas ID from the URL
	KelasID := c.Param("id")

	var kelas models.Kelas
	if err := database.DB.Where("id = ?", KelasID).First(&kelas).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kelas not found"})
		return
	}

	// Bind the input JSON
	var input UpdateKelasInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.NamaKelas != nil {
		kelas.NamaKelas = *input.NamaKelas
	}

	// Save the updated Kelas
	if err := database.DB.Save(&kelas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Kelas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kelas updated successfully", "Kelas": kelas})
}

func DeleteKelas(c *gin.Context) {
	// Get the Kelas ID from the URL
	KelasID := c.Param("id")

	var kelas models.Kelas
	if err := database.DB.Where("id = ?", KelasID).First(&kelas).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kelas not found"})
		return
	}

	// Delete the Kelas
	if err := database.DB.Delete(&kelas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Kelas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kelas deleted successfully"})
}
