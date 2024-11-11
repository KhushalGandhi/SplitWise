package migrations

import (
	"gorm.io/gorm"
	"splitwise/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Spend{}, &models.Log{}, &models.Group{}, &models.Account{})
}
