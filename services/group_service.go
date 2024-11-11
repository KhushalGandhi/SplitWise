package services

import (
	"errors"
	"gorm.io/gorm"
	"splitwise/models"
	"splitwise/repositories"
)

func CreateGroup(group *models.GroupRequest) error {
	baseModel := models.Group{
		Model:     gorm.Model{},
		Name:      group.Name,
		CreatorID: group.CreatorID,
	}
	return repositories.CreateGroup(&baseModel)
}

func GetGroupDetails(groupID uint) (models.Group, error) {
	return repositories.GetGroupByID(groupID)
}

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

func AddUserToGroup(groupID, newUserID uint) error {
	return repositories.AddUserToGroup(groupID, newUserID)
}
