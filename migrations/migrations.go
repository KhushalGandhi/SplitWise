package migrations

import (
	"gorm.io/gorm"
	"splitwise/models"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Spend{}, &models.Log{}, &models.Group{}, &models.Account{}, &models.User{}, &models.Share{})
	if err != nil {
		return
	}
}
