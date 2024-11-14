package services

import (
	"errors"
	"splitwise/models"
	"splitwise/repositories"
)

func CreateSpend(req models.CreateSpendRequest) error {

	// Retrieve the email associated with the provided user ID
	email, err := repositories.GetEmailByUserID(req.UserID)
	if err != nil {
		return err
	}

	// Check if the email belongs to a user in the specified group
	isMember, err := repositories.IsEmailInGroup(email, req.GroupID)
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("user is not a member of the specified group")
	}

	// Initialize the spend struct
	spend := models.Spend{
		GroupID:     req.GroupID,
		UserID:      req.UserID,
		Amount:      req.Amount,
		Description: req.Description,
		SplitType:   req.SplitType,
		Status:      "pending",
	}

	// Map to hold calculated shares for each user's email in the group
	splits := make(map[string]float64)

	// Calculate shares based on the split type
	if req.SplitType == "equal" {

		// this is one of my doubt that if we are not sure about the parts but can be checked through
		// frontend if we provide them with one more api that will tell how many members are there so that condition/check can be added there

		// Get the total number of members in the group
		numMembers, err := repositories.GetGroupMemberCount(req.GroupID)
		if err != nil {
			return err
		}

		// Calculate the equal share per member
		perPersonAmount := req.Amount / float64(numMembers)
		for email := range req.SplitValues {
			splits[email] = perPersonAmount
		}
	} else if req.SplitType == "exact" {
		total := 0.0
		for email, amount := range req.SplitValues {
			splits[email] = amount
			total += amount
		}

		// Validate that exact splits total the given amount
		if total != req.Amount {
			return errors.New("exact split values do not sum to the total amount")
		}
	} else if req.SplitType == "percentage" {
		total := 0.0
		for email, percent := range req.SplitValues {
			amount := (percent / 100) * req.Amount
			splits[email] = amount
			total += amount
		}

		// Validate that the total of percentage splits matches the requested amount
		if total != req.Amount {
			return errors.New("percentage split does not sum to the total amount")
		}
	} else {
		return errors.New("invalid split type")
	}

	// Save the spend entry to the database
	if err := repositories.CreateSpend(&spend); err != nil {
		return err
	}

	// Helper function to get userID from email
	getUserID := func(email string) (uint, error) {
		return repositories.GetAccountIdByEmail(email)
	}

	// Save each share based on calculated splits
	for email, amount := range splits {
		userID, err := getUserID(email)
		if err != nil {
			return err
		}
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

//
//func CreateSpend(req models.CreateSpendRequest) error {
//
//	email, err := repositories.GetEmailByUserID(req.UserID)
//	if err != nil {
//		return err
//	}
//
//	// Check if the email belongs to a user in the specified group
//	isMember, err := repositories.IsEmailInGroup(email, req.GroupID)
//	if err != nil {
//		return err
//	}
//	if !isMember {
//		return errors.New("user is not a member of the specified group")
//	}
//
//	spend := models.Spend{
//		GroupID:     req.GroupID,
//		UserID:      req.UserID,
//		Amount:      req.Amount,
//		Description: req.Description,
//		SplitType:   req.SplitType,
//		Status:      "pending",
//	}
//
//	// Determine shares based on split type
//	splits := make(map[uint]float64)
//	if req.SplitType == "equal" {
//		numMembers, err := repositories.GetGroupMemberCount(req.GroupID)
//		if err != nil {
//			return err
//		}
//		perPersonAmount := req.Amount / float64(numMembers)
//		for userID := range req.SplitValues {
//			splits[userID] = perPersonAmount
//		}
//	} else if req.SplitType == "exact" {
//		total := 0.0
//		for userID, amount := range req.SplitValues {
//			splits[userID] = amount
//			total += amount
//		}
//		if total != req.Amount {
//			return errors.New("exact split values do not sum to the total amount")
//		}
//	} else if req.SplitType == "percentage" {
//		total := 0.0
//		for userID, percent := range req.SplitValues {
//			amount := (percent / 100) * req.Amount
//			splits[userID] = amount
//			total += amount
//		}
//		if total != req.Amount {
//			return errors.New("percentage split does not sum to the total amount")
//		}
//	} else {
//		return errors.New("invalid split type")
//	}
//
//	// Save the spend
//	if err := repositories.CreateSpend(&spend); err != nil {
//		return err
//	}
//
//	// Save each share based on calculated splits
//	for userID, amount := range splits {
//		share := models.Share{
//			SpendID: spend.ID,
//			UserID:  userID,
//			Amount:  amount,
//		}
//		if err := repositories.CreateShare(&share); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
