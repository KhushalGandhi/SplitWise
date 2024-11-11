package services

import (
	"gorm.io/gorm"
	"splitwise/models"
	"splitwise/repositories"
)

func CreateSpend(spend *models.SpendRequest) error {
	baseModel := models.Spend{
		Model:     gorm.Model{},
		GroupID:   spend.GroupID,
		Amount:    spend.Amount,
		SpenderID: spend.SpenderID,
		SplitType: spend.SplitType,
	}
	return repositories.CreateSpend(&baseModel)
}

func GetGroupSpends(groupID uint) ([]models.Spend, error) {
	return repositories.GetSpendsByGroupID(groupID)
}
