package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"` // "admin", "instructor", "student"
}
