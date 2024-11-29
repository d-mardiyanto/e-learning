package controllers

import (
	"e-learning/database"
	"e-learning/models"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CourseInput struct {
	KodeMateri string `json:"kode_materi" binding:"required"`
	Judul      string `json:"judul"`
	Kelas      string `json:"kelas"`
	Prodi      string `json:"prodi"`
}

type UpdateCourseInput struct {
	Judul *string `json:"judul"`
	Kelas *string `json:"kelas"`
	Prodi *string `json:"prodi"`
}

func ShowCourses(c *gin.Context) {
	var courses []models.Course

	// Fetch all courses from the database
	if err := database.DB.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses", "details": err.Error()})
		return
	}

	// Respond with the list of courses
	c.JSON(http.StatusOK, gin.H{
		"message": "Courses retrieved successfully",
		"courses": courses,
	})
}

func GetCourse(c *gin.Context) {
	courseID := c.Param("id")
	id, err := strconv.Atoi(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var course models.Course
	if err := database.DB.Preload("CourseFiles").First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Course retrieved successfully",
		"course":  course,
	})
}

func CreateCourse(c *gin.Context) {
	instructorID := c.MustGet("userID").(uint)

	// Bind and validate input
	var input CourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new course
	course := models.Course{
		KodeMateri: input.KodeMateri,
		Judul:      input.Judul,
		Kelas:      input.Kelas,
		Prodi:      input.Prodi,
		CreatedBy:  instructorID,
	}

	if err := database.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create course"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Course created successfully",
		"course":  course,
	})
}

func UpdateCourse(c *gin.Context) {
	// Get the course ID from the URL
	courseID := c.Param("id")

	// Check if the course exists and belongs to the instructor
	instructorID := c.MustGet("userID").(uint)
	var course models.Course
	if err := database.DB.Where("id = ? AND createdBy = ?", courseID, instructorID).First(&course).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Bind the input JSON
	var input UpdateCourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Judul != nil {
		course.Judul = *input.Judul
	}
	if input.Kelas != nil {
		course.Kelas = *input.Kelas
	}
	if input.Prodi != nil {
		course.Prodi = *input.Prodi
	}

	// Save the updated course
	if err := database.DB.Save(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully", "course": course})
}

func DeleteCourse(c *gin.Context) {
	// Get the course ID from the URL
	courseID := c.Param("id")

	// Check if the course exists and belongs to the instructor
	instructorID := c.MustGet("userID").(uint)
	var course models.Course
	if err := database.DB.Where("id = ? AND instructor_id = ?", courseID, instructorID).First(&course).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Delete the course
	if err := database.DB.Delete(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

// UploadCourseFile handles the file upload for a course material
func UploadCourseFile(c *gin.Context) {
	// Parse form data
	idMateri := c.PostForm("id_materi")
	fileType := c.PostForm("file_type")
	fileLabel := c.PostForm("file_label")
	orderNumberStr := c.PostForm("order_number") // Order number as string

	// Validate required fields
	if idMateri == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_materi is required"})
		return
	}

	// Convert orderNumberStr to int
	orderNumber, err := strconv.Atoi(orderNumberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_number; it must be an integer"})
		return
	}

	// Handle file upload
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed", "details": err.Error()})
		return
	}

	// Save the file to a local directory
	uploadPath := "uploads/courses/" + idMateri
	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
	filePath := filepath.Join(uploadPath, fileName)

	// Create the directory if it doesn't exist
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "details": err.Error()})
		return
	}

	// Create a new CourseFiles record in the database
	courseFile := models.CourseFiles{
		IdMateri:    idMateri,
		FileType:    fileType,
		FileLabel:   fileLabel,
		OrderNumber: orderNumber, // Assign the converted int value
		File:        filePath,
	}

	if err := database.DB.Create(&courseFile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file record", "details": err.Error()})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file":    courseFile,
	})
}
