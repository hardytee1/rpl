package models

import(
	"gorm.io/gorm"
)

type Wts struct {
	gorm.Model
	Mata_pelajaran string `json:"mata_pelajaran" binding:"required"`
	TeacherID uint
}