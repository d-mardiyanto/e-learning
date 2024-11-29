package models

import "github.com/jinzhu/gorm"

type Prodi struct {
	gorm.Model
	IDProdi   string `gorm:"primaryKey" json:"id" binding:"required"`
	NamaProdi string `json:"nama_prodi" binding:"required"`
}
