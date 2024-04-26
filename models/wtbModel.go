package models

import(
	"gorm.io/gorm"
)

type Wtb struct {
	gorm.Model
	Mata_pelajaran string `json:"mata_pelajaran" binding:"required"`
	Jumlah_pertemuan string `json:"jumlah_pertemuan" binding:"required"`
	Harga string `json:"harga" binding:"required"`
	UserID uint
}