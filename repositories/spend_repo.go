package repositories

import (
	"splitwise/db"
	"splitwise/models"
)

func CreateSpend(spend *models.Spend) error {
	return database.DB.Create(spend).Error
}

func GetSpendsByGroupID(groupID uint) ([]models.Spend, error) {
	var spends []models.Spend
	err := database.DB.Where("group_id = ?", groupID).Find(&spends).Error
	return spends, err
}
