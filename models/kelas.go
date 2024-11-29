package models

import "github.com/jinzhu/gorm"

type Kelas struct {
	gorm.Model
	IDKelas   string `gorm:"primaryKey" json:"id" binding:"required"`
	NamaKelas string `json:"nama_kelas" binding:"required"`
}
