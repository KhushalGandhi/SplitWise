package services

import (
	"fmt"
	"splitwise/repositories"
)

func CalculateRemainingBalance(groupID string) (float64, error) {
	spends, err := repositories.GetSpendsByGroupID(groupID)
	fmt.Println(1)
	if err != nil {
		return 0, err
	}

	fmt.Println(2)
	fmt.Println(spends)

	var totalBalance float64
	for _, spend := range spends {
		totalBalance += spend.Amount
	}
	return totalBalance, nil
}
