package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	UserID         int             `json:"user_id" binding:"required"`
	QuizID         int             `json:"quiz_id"  binding:"required"`
	CompletionDate string          `gorm:"type:date" json:"completion_date"  binding:"required"`
	StartTime      time.Time       `gorm:"column:start_time;type:time" json:"start_time" binding:"required"`
	EndTime        time.Time       `gorm:"column:end_time;type:time" json:"end_time"`
	CreatedAt      time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time       `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	Detail         []Answer_Detail `gorm:"foreignKey:answer_id" json:"answers"` // Association
}

type Answer_Detail struct {
	gorm.Model
	AnswerID   int       `json:"answer_id" gorm:"constraint:OnDelete:CASCADE;" binding:"required"`
	QuestionID int       `json:"question_id" binding:"required"`
	Answer     string    `json:"answer"  binding:"required"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
