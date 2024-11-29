package models

import "github.com/jinzhu/gorm"

type Students struct {
	gorm.Model
	Nama        string              `json:"nama" binding:"required"`
	Gender      string              `json:"gender"`
	Birthdate   string              `json:"birthdate"`
	Birthplace  string              `json:"birthplace"`
	Address     string              `json:"address"`
	Phone       string              `json:"phone"`
	Email       string              `json:"email"`
	EnteredYear int                 `json:"entered_year"`
	Academic    []Students_Academic `gorm:"foreignKey:id_student" json:"course_files"` // Association
}

type Students_Academic struct {
	gorm.Model
	IDStudent string `json:"id_materi" binding:"required"`
	Semester  string `json:"semester"`
	Kelas     string `json:"kelas"`
	Prodi     string `json:"prodi"`
}
