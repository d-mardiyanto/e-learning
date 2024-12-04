package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Students struct {
	gorm.Model
	Name        string              `json:"name" binding:"required"`
	Gender      string              `json:"gender" binding:"required"`
	Birthdate   string              `json:"birthdate" binding:"required"`
	Birthplace  string              `json:"birthplace"`
	Address     string              `json:"address"`
	Phone       string              `json:"phone" binding:"required"`
	Email       string              `json:"email" binding:"required"`
	EnteredYear int                 `gorm:"type:year" json:"entered_year"`
	CreatedAt   time.Time           `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time           `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	Academic    []Students_Academic `gorm:"foreignKey:student_id" json:"academic"` // Association
}

type Students_Academic struct {
	gorm.Model
	StudentId    string    `json:"student_id" gorm:"constraint:OnDelete:CASCADE;" binding:"required"`
	Semester     string    `json:"semester"`
	Year         int       `gorm:"type:year" json:"year"`
	StartDate    time.Time `gorm:"type:date" json:"start_date"`
	EndDate      time.Time `gorm:"type:date" json:"end_date"`
	Classes      string    `json:"classes" binding:"required"`
	StudyProgram string    `json:"program_study" binding:"required"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
