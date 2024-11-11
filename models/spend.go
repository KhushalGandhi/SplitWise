package models

import "gorm.io/gorm"

//type Spend struct {
//	gorm.Model
//	GroupID   uint    `json:"group_id"`
//	Amount    float64 `json:"amount"`
//	SpenderID uint    `json:"spender_id"`
//	SplitType string  `json:"split_type"`
//}

type CreateSpendRequest struct {
	GroupID     uint             `json:"group_id"`
	UserID      uint             `json:"user_id"`
	Amount      float64          `json:"amount"`
	Description string           `json:"description"`
	SplitType   string           `json:"split_type"`   // "equal", "exact", "percentage"
	SplitValues map[uint]float64 `json:"split_values"` // Key: UserID, Value: Split Amount/Percentage
}

type Spend struct {
	gorm.Model
	GroupID     uint    `json:"group_id"`
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	SplitType   string  `json:"split_type"` // "equal", "exact", "percentage"
	Status      string  `json:"status"`     // "pending" or "settled"
	Shares      []Share `gorm:"foreignKey:SpendID"`
}

type Share struct {
	gorm.Model
	SpendID uint    `json:"spend_id"`
	UserID  uint    `json:"user_id"`
	Amount  float64 `json:"amount"`
}
