package controllers

import (
	"e-learning/database"
	"e-learning/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentsInput struct {
	Nama        string `json:"nama" binding:"required"`
	Gender      string `json:"gender"`
	Birthdate   string `json:"birthdate"`
	Birthplace  string `json:"birthplace"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	EnteredYear int    `json:"entered_year"`
}

type UpdateStudentsInput struct {
	Nama        *string `json:"nama" binding:"required"`
	Gender      *string `json:"gender"`
	Birthdate   *string `json:"birthdate"`
	Birthplace  *string `json:"birthplace"`
	Address     *string `json:"address"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	EnteredYear *int    `json:"entered_year"`
}

func ShowStudents(c *gin.Context) {
	var students []models.Students

	// Fetch all Students from the database
	if err := database.DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Students", "details": err.Error()})
		return
	}

	// Respond with the list of Studentss
	c.JSON(http.StatusOK, gin.H{
		"message":  "Students retrieved successfully",
		"students": students,
	})
}

func GetStudent(c *gin.Context) {
	StudentsID := c.Param("id")
	id, err := strconv.Atoi(StudentsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Students ID"})
		return
	}

	var students models.Students
	if err := database.DB.Preload("Students_Academic").First(&students, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Students not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Students retrieved successfully",
		"students": students,
	})
}

func CreateStudent(c *gin.Context) {
	// Bind and validate input
	var input StudentsInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	students := models.Students{
		Nama:        input.Nama,
		Gender:      input.Gender,
		Birthdate:   input.Birthdate,
		Birthplace:  input.Birthplace,
		Address:     input.Address,
		Phone:       input.Phone,
		Email:       input.Email,
		EnteredYear: input.EnteredYear,
	}

	if err := database.DB.Create(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Students"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Students created successfully",
		"students": students,
	})
}

func UpdateStudent(c *gin.Context) {
	// Get the Students ID from the URL
	StudentsID := c.Param("id")

	var students models.Students
	if err := database.DB.Where("id = ?", StudentsID).First(&students).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Students not found"})
		return
	}

	// Bind the input JSON
	var input UpdateStudentsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated Students
	if err := database.DB.Save(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Students"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Students updated successfully", "students": students})
}

func DeleteStudent(c *gin.Context) {
	// Get the Students ID from the URL
	StudentsID := c.Param("id")

	var students models.Students
	if err := database.DB.Where("id = ?", StudentsID).First(&students).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Students not found"})
		return
	}

	// Delete the Students
	if err := database.DB.Delete(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Students"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Students deleted successfully"})
}
