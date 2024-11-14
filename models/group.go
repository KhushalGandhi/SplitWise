package models

import (
	"time"
)

type Group struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatorID uint      `gorm:"creator_id" json:"creator_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
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
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GroupID   string    `gorm:"group_id" json:"group_id"`
	Name      string    `gorm:"name" json:"name"`
	Email     string    `gorm:"email" json:"email" gorm:"unique"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
