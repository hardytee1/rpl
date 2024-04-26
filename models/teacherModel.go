package models

import(
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Pendidikan string `json:"pendidikan" binding:"required"`
	Bukti_1 string `json:"bukti_1" binding:"required"`
	UserID uint `gorm:"unique"`
	Wts []Wts
}