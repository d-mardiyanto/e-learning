package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Classes struct {
	gorm.Model
	ClassName string    `json:"class_name" binding:"required"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
