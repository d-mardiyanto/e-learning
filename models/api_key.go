package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type APIKey struct {
	gorm.Model
	Token     string    `gorm:"uniqueIndex;not null" json:"token"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
