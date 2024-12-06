package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Instructors struct {
	gorm.Model
	Photo      string    `json:"photo"`
	Name       string    `json:"name" binding:"required"`
	Profession string    `json:"profession" binding:"required"`
	Email      string    `json:"email" binding:"required"`
	Phone      string    `json:"phone" binding:"required"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
