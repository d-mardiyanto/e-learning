package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type APIKey struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"uniqueIndex;not null" json:"key_value"`
	CreatedAt time.Time `json:"created_at"`
}
