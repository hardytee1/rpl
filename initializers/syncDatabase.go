package initializers

import "github.com/hardytee1/rpl/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}