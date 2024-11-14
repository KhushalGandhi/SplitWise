package repositories

import (
	"fmt"
	"splitwise/db"
	"splitwise/models"
)

func CreateSpend(spend *models.Spend) error {
	return database.DB.Create(spend).Error
}

func GetSpendsByGroupID(groupID string) ([]models.Spend, error) {
	var spends []models.Spend
	err := database.DB.Where("group_id = ?", groupID).Find(&spends).Error
	if err != nil {
		fmt.Println("Error fetching spends:", err)
	}
	return spends, err
}

func CreateShare(share *models.Share) error {
	return database.DB.Create(&share).Error
}
