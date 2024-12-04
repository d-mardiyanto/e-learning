package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Quiz struct {
	gorm.Model
	CourseID  string      `json:"course_id" binding:"required"`
	PassMark  int         `json:"pass_mark" binding:"required"`
	TimeLimit int         `json:"time_limit" binding:"required"`
	Attempt   int         `json:"attempt" binding:"required"`
	QuizName  string      `json:"quiz_name" binding:"required"`
	StartDate string      `gorm:"type:date" json:"start_date" binding:"required"`
	EndDate   string      `gorm:"type:date" json:"end_date" binding:"required"`
	CreatedAt time.Time   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	Questions []Questions `gorm:"foreignKey:quiz_id" json:"pertanyaan"` // Association
}

type Questions struct {
	gorm.Model
	QuizID        string    `json:"quiz_id" gorm:"constraint:OnDelete:CASCADE;" binding:"required"`
	QuestionType  string    `json:"question_type" binding:"required"`
	Question      string    `json:"question" binding:"required"`
	Answer        string    `json:"answer" binding:"required"`
	CorrectAnswer string    `json:"correct_answer" binding:"required"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
