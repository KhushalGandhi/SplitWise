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
