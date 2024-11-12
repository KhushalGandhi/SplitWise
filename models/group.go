package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name      string `json:"name"`
	CreatorID uint   `json:"creator_id"`
}

type GroupRequest struct {
	Name      string `json:"name"`
	CreatorID uint   `json:"creator_id"`
}

type Name struct {
	Name string `json:"name"`
}

type GroupMember struct {
	GroupId uint `json:"group_id"`
	UserId  uint `json:"user_id"`
}

type AddUserToGroupRequest struct {
	GroupID string `json:"group_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type User struct {
	gorm.Model
	GroupID string `json:"group_id"`
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
}
