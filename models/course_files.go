package models

import "github.com/jinzhu/gorm"

type CourseFiles struct {
	gorm.Model
	IdMateri    string `json:"id_materi" binding:"required"`
	FileType    string `json:"file_type"`
	FileLabel   string `json:"file_label"`
	OrderNumber int    `json:"order_number"`
	File        string `json:"file"`
}
