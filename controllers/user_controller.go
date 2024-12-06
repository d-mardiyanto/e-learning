package controllers

import (
	"e-learning/database"
	"e-learning/models"
	"e-learning/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var input RegisterInput

	// Bind JSON input and validate it
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the struct
	if err := utils.ValidateStruct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, _ := utils.HashPassword(input.Password)

	// Create a new user
	user := models.User{Name: input.Name, Email: input.Email, Password: hashedPassword, Username: input.Username, Role: "student"}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func GetUserProfile(c *gin.Context) {
	// Get user ID from the context (set by AuthMiddleware)
	userID := c.MustGet("userID").(uint)

	// Fetch user from the database
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}
