package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Title        string         `json:"title" binding:"required"`
	Thumbnail    string         `json:"thumbnail"`
	Description  string         `json:"description"`
	Classes      string         `json:"classes" binding:"required"`
	StudyProgram string         `json:"program_study" binding:"required"`
	CreatedBy    int            `json:"created_by" binding:"required"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	CourseFiles  []CourseFiles  `gorm:"foreignKey:course_id" json:"course_files"`  // Association
	Rating       []CourseRating `gorm:"foreignKey:course_id" json:"course_rating"` // Association
}

type CourseFiles struct {
	gorm.Model
	CourseID    string    `json:"course_id" gorm:"constraint:OnDelete:CASCADE;" binding:"required"`
	FileType    string    `json:"file_type" binding:"required"`
	FileLabel   string    `json:"file_label" binding:"required"`
	OrderNumber int       `json:"order_number" binding:"required"`
	File        string    `json:"file" binding:"required"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}

type CourseRating struct {
	gorm.Model
	CourseID  string    `json:"course_id" gorm:"constraint:OnDelete:CASCADE;" binding:"required"`
	Rate      string    `json:"rate" binding:"required"`
	RatedBy   string    `json:"rated_by"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
