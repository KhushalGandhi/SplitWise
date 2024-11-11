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
