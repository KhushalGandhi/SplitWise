package models

import "gorm.io/gorm"

type Spend struct {
	gorm.Model
	GroupID   uint    `json:"group_id"`
	Amount    float64 `json:"amount"`
	SpenderID uint    `json:"spender_id"`
	SplitType string  `json:"split_type"`
}

type SpendRequest struct {
	GroupID   uint    `json:"group_id"`
	Amount    float64 `json:"amount"`
	SpenderID uint    `json:"spender_id"`
	SplitType string  `json:"split_type"`
}
