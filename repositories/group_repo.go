package repositories

import (
	"errors"
	"gorm.io/gorm"
	"splitwise/db"
	"splitwise/models"
)

func CreateGroup(group *models.Group) error {
	return database.DB.Create(group).Error
}

func GetGroupByID(id uint) (models.Group, error) {
	var group models.Group
	err := database.DB.First(&group, id).Error
	return group, err
}

func GetUnsettledSpendsByGroupID(groupID uint) ([]models.Spend, error) {
	var spends []models.Spend
	if err := database.DB.Where("group_id = ? AND status = ?", groupID, "pending").Find(&spends).Error; err != nil {
		return nil, err
	}
	return spends, nil
}

func GetGroupMemberCount(groupID string) (int, error) {
	var count int64
	if err := database.DB.Model(&models.User{}).
		Where("group_id = ?", groupID).
		Distinct("email").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

//
//func AddUserToGroup(groupID, newUserID uint) error {
//	member := models.GroupMember{
//		GroupId: groupID,
//		UserId:  newUserID,
//	}
//	return database.DB.Create(&member).Error
//}

func DeleteGroup(group *models.Group) error {
	return database.DB.Delete(group).Error
}

func IsGroupCreator(groupID string, userID uint) (bool, error) {
	var group models.Group
	err := database.DB.Where("group_id = ? AND creator_id = ?", groupID, userID).First(&group).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// AddUserToGroup inserts a new user record into the group.
func AddUserToGroup(user models.User) error {
	// Check if the user already exists in the group by email
	var existingUser models.User
	err := database.DB.Where("group_id = ? AND email = ?", user.GroupID, user.Email).First(&existingUser).Error
	if err == nil {
		return errors.New("user already exists in the group")
	}

	// Add user to the group
	return database.DB.Create(&user).Error
}

func IsUserInGroup(userID uint, groupID string) (bool, error) {
	var user models.User
	if err := database.DB.Where("id = ? AND group_id = ?", userID, groupID).First(&user).Error; err != nil {
		// If record is not found, return false
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func IsEmailInGroup(email string, groupID string) (bool, error) {
	var user models.User
	if err := database.DB.Where("email = ? AND group_id = ?", email, groupID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
