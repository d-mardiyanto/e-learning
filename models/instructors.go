package models

import "github.com/jinzhu/gorm"

type Instructors struct {
	gorm.Model
	Nama  string `json:"name" binding:"required"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
