package services

import (
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
