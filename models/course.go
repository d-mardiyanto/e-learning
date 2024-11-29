package models

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	KodeMateri  string        `json:"kode_materi" binding:"required"`
	Judul       string        `json:"judul"`
	Kelas       string        `json:"kelas"`
	Prodi       string        `json:"prodi"`
	CreatedBy   uint          `json:"created_by"`
	CourseFiles []CourseFiles `gorm:"foreignKey:id_materi" json:"course_files"` // Association
}
