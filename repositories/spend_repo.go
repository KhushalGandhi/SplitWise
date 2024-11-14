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

//func GetSpendsByGroupIDandUserId(groupID string, userId uint) ([]models.Spend, error) {
//	var spends []models.Spend
//	err := database.DB.Where("group_id = ? and user_id  = ?", groupID, userId).Find(&spends).Error
//	if err != nil {
//		fmt.Println("Error fetching spends:", err)
//	}
//	return spends, err
//}

// Fetches all spend IDs for a given group ID

func GetSpendIDsByGroupID(groupID string) ([]uint, error) {
	var spends []models.Spend
	var spendIDs []uint

	err := database.DB.Where("group_id = ?", groupID).Find(&spends).Error
	if err != nil {
		return nil, err
	}

	for _, spend := range spends {
		spendIDs = append(spendIDs, spend.ID)
	}
	return spendIDs, nil
}

// Fetches all shares for a given list of spend IDs and user ID

func GetSharesBySpendIDsAndUserID(spendIDs []uint, userID uint) ([]models.Share, error) {
	var shares []models.Share
	err := database.DB.Where("spend_id IN ? AND user_id = ?", spendIDs, userID).Find(&shares).Error
	return shares, err
}

func CreateShare(share *models.Share) error {
	return database.DB.Create(&share).Error
}
