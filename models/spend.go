package models

import "time"

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
	ID          uint    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GroupID     uint    `gorm:"group_id" json:"group_id"`
	UserID      uint    `gorm:"user_id" json:"user_id"`
	Amount      float64 `gorm:"amount" json:"amount"`
	Description string  `gorm:"description" json:"description"`
	SplitType   string  `gorm:"split_type" json:"split_type"` // "equal", "exact", "percentage"
	Status      string  `gorm:"status" json:"status"`         // "pending" or "settled"
	Shares      []Share `gorm:"shares" gorm:"foreignKey:SpendID"`
}

type Share struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SpendID   uint      `json:"spend_id"`
	UserID    uint      `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
