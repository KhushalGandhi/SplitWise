package repositories

import (
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

func GetGroupMemberCount(groupID uint) (int, error) {
	var count int64
	if err := database.DB.Model(&models.GroupMember{}).Where("group_id = ?", groupID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func AddUserToGroup(groupID, newUserID uint) error {
	member := models.GroupMember{
		GroupId: groupID,
		UserId:  newUserID,
	}
	return database.DB.Create(&member).Error
}

func DeleteGroup(group *models.Group) error {
	return database.DB.Delete(group).Error
}
