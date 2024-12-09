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
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description"`
	Thumbnail    string `json:"thumbnail"`
	CreatedBy    int    `json:"created_by" binding:"required"`
	Classes      string `json:"classes" binding:"required"`
	StudyProgram string `json:"program_study" binding:"required"`
}

type UpdateCourseInput struct {
	Title        *string `json:"title" binding:"required"`
	Description  *string `json:"description"`
	Thumbnail    *string `json:"thumbnail"`
	Classes      *string `json:"classes" binding:"required"`
	StudyProgram *string `json:"program_study" binding:"required"`
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
	// Bind and validate input
	var input CourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new course
	course := models.Course{
		Title:        input.Title,
		Description:  input.Description,
		Classes:      input.Classes,
		StudyProgram: input.StudyProgram,
		CreatedBy:    input.CreatedBy,
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
	var course models.Course
	if err := database.DB.Where("id = ?", courseID).First(&course).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Bind the input JSON
	var input CourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new course
	course.Title = input.Title
	course.Description = input.Description
	course.Classes = input.Classes
	course.StudyProgram = input.StudyProgram

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

	var course models.Course
	if err := database.DB.Where("id = ?", courseID).First(&course).Error; err != nil {
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
	CourseID := c.PostForm("course_id")
	fileType := c.PostForm("file_type")
	fileLabel := c.PostForm("file_label")
	orderNumberStr := c.PostForm("order_number") // Order number as string

	// Validate required fields
	if CourseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course ID is required"})
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
	uploadPath := "uploads/courses/" + CourseID
	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
	filePath := filepath.Join(uploadPath, fileName)

	// Create the directory if it doesn't exist
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "details": err.Error()})
		return
	}

	// Create a new CourseFiles record in the database
	courseFile := models.CourseFiles{
		CourseID:    CourseID,
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

// UploadCourseFile handles the file upload for a course material
func UploadThumbnail(c *gin.Context) {
	// Parse form data
	CourseID := c.PostForm("course_id")

	// Validate required fields
	if CourseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course ID is required"})
		return
	}

	// Handle file upload
	file, err := c.FormFile("thumbnail")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thumbnail upload failed", "details": err.Error()})
		return
	}

	// Save the file to a local directory
	uploadPath := "uploads/courses/" + CourseID + "/thumbnail"
	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
	filePath := filepath.Join(uploadPath, fileName)

	// Create the directory if it doesn't exist
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "details": err.Error()})
		return
	}

	// Create a new CourseFiles record in the database
	var course models.Course
	course.Thumbnail = filePath

	if err := database.DB.Model(&models.Course{}).Where("id = ?", CourseID).UpdateColumn("thumbnail", filePath).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file record", "details": err.Error()})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file":    course,
	})
}
