package services

import (
	"errors"
	"splitwise/models"
	"splitwise/repositories"
	"splitwise/utils"
	"time"
)

func CreateGroup(group *models.GroupRequest) error {
	baseModel := models.Group{
		Name:      group.Name,
		CreatorID: group.CreatorID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	if err := repositories.CreateGroup(&baseModel); err != nil {
		return err
	}

	creatorAccount, err := repositories.GetAccountByID(group.CreatorID)
	if err != nil {
		return err
	}

	user := models.User{
		GroupID:   utils.GenerateGroupID(),
		Name:      group.Name,
		Email:     creatorAccount.Email,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	return repositories.AddUserToGroup(user)
}

//func GetGroupDetails(groupID uint) (models.Group, error) {
//	return repositories.GetGroupByID(groupID)
//}

func CanDeleteGroup(groupID, userID uint) error {
	group, err := repositories.GetGroupByID(groupID)
	if err != nil {
		return err
	}

	if group.CreatorID != userID {
		return errors.New("only the group creator can delete this group")
	}

	unsettledSpends, err := repositories.GetUnsettledSpendsByGroupID(groupID)
	if err != nil {
		return err
	}
	if len(unsettledSpends) > 0 {
		return errors.New("cannot delete group with unsettled spends")
	}

	return nil
}

func CanAddUserToGroup(groupID, userID uint) error {
	group, err := repositories.GetGroupByID(groupID)
	if err != nil {
		return err
	}

	if group.CreatorID != userID {
		return errors.New("only the group creator can add users")
	}

	return nil
}

func DeleteGroup(groupID, userID uint) error {
	// Ensure only the creator can delete and all spends are settled
	if err := CanDeleteGroup(groupID, userID); err != nil {
		return err
	}

	group, err := repositories.GetGroupByID(groupID)
	if err != nil {
		return err
	}

	return repositories.DeleteGroup(&group)
}

func AddUserToGroup(creatorID string, req models.AddUserToGroupRequest) error {
	// Check if the user is the creator of the group
	isCreator, err := repositories.IsGroupCreator(req.GroupID, creatorID)
	if err != nil {
		return err
	}
	if !isCreator {
		return errors.New("only the group creator can add users to the group")
	}

	// Prepare user data to add to group
	user := models.User{
		GroupID: req.GroupID,
		Name:    req.Name,
		Email:   req.Email,
	}

	// Add user to the group via repository
	return repositories.AddUserToGroup(user)
}
