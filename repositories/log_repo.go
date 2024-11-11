package repositories

import (
	"splitwise/db"
	"splitwise/models"
)

func CreateLog(log *models.Log) error {
	return database.DB.Create(log).Error
}
