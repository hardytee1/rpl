package models

import(
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
	Nama string `json:"nama" binding:"required"`
	Biodata string `json:"biodata" binding:"required"`
	Notelpon string `json:"notelpon" binding:"required"`
	Role Role `gorm:"default: user"`
}