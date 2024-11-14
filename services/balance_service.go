package services

import (
	"splitwise/repositories"
)

func CalculateRemainingBalance(groupID string) (float64, error) {
	spends, err := repositories.GetSpendsByGroupID(groupID)
	//fmt.Println(1)
	if err != nil {
		return 0, err
	}

	//fmt.Println(2)
	//fmt.Println(spends)

	var totalBalance float64
	for _, spend := range spends {
		totalBalance += spend.Amount
	}
	return totalBalance, nil
}

func CalculateRemainingBalanceforUser(groupID string, userID uint) (float64, error) {
	// Step 1: Get all spend IDs for the group
	spendIDs, err := repositories.GetSpendIDsByGroupID(groupID)
	if err != nil {
		return 0, err
	}

	// Step 2: Get all shares for the user within the group based on spend IDs
	shares, err := repositories.GetSharesBySpendIDsAndUserID(spendIDs, userID)
	if err != nil {
		return 0, err
	}

	// Step 3: Calculate the total balance based on shares
	var totalBalance float64
	for _, share := range shares {
		totalBalance += share.Amount
	}
	return totalBalance, nil
}
