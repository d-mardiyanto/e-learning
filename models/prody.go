package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type StudyProgram struct {
	gorm.Model
	ProgramName string    `json:"program_name" binding:"required"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
