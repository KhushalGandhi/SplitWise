package services

import (
	"errors"
	"splitwise/models"
	"splitwise/repositories"
)

func CreateSpend(req *models.CreateSpendRequest) error {
	spend := models.Spend{
		GroupID:     req.GroupID,
		UserID:      req.UserID,
		Amount:      req.Amount,
		Description: req.Description,
		SplitType:   req.SplitType,
		Status:      "pending",
	}

	// Determine shares based on split type
	splits := make(map[uint]float64)
	if req.SplitType == "equal" {
		numMembers, err := repositories.GetGroupMemberCount(req.GroupID)
		if err != nil {
			return err
		}
		perPersonAmount := req.Amount / float64(numMembers)
		for userID := range req.SplitValues {
			splits[userID] = perPersonAmount
		}
	} else if req.SplitType == "exact" {
		total := 0.0
		for userID, amount := range req.SplitValues {
			splits[userID] = amount
			total += amount
		}
		if total != req.Amount {
			return errors.New("exact split values do not sum to the total amount")
		}
	} else if req.SplitType == "percentage" {
		total := 0.0
		for userID, percent := range req.SplitValues {
			amount := (percent / 100) * req.Amount
			splits[userID] = amount
			total += amount
		}
		if total != req.Amount {
			return errors.New("percentage split does not sum to the total amount")
		}
	} else {
		return errors.New("invalid split type")
	}

	// Save the spend
	if err := repositories.CreateSpend(&spend); err != nil {
		return err
	}

	// Save each share based on calculated splits
	for userID, amount := range splits {
		share := models.Share{
			SpendID: spend.ID,
			UserID:  userID,
			Amount:  amount,
		}
		if err := repositories.CreateShare(&share); err != nil {
			return err
		}
	}

	return nil
}

//func GetGroupSpends(groupID uint) ([]models.Spend, error) {
//	return repositories.GetSpendsByGroupID(groupID)
//}
