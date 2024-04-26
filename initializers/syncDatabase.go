package initializers

import "github.com/hardytee1/rpl/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Teacher{})
	DB.AutoMigrate(&models.Wts{})
	DB.AutoMigrate(&models.Wtb{})
}