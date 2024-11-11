package services

import (
	"splitwise/repositories"
)

func CalculateRemainingBalance(groupID uint) (float64, error) {
	spends, err := repositories.GetSpendsByGroupID(groupID)
	if err != nil {
		return 0, err
	}

	var totalBalance float64
	for _, spend := range spends {
		totalBalance += spend.Amount
	}
	return totalBalance, nil
}
