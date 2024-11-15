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
	GroupID     string             `json:"group_id"`
	UserID      uint               `json:"user_id"`
	Amount      float64            `json:"amount"`
	Description string             `json:"description"`
	SplitType   string             `json:"split_type"`   // "equal", "exact", "percentage"
	SplitValues map[string]float64 `json:"split_values"` // Key: UserID, Value: Split Amount/Percentage
}

type Spend struct {
	ID          uint    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GroupID     string  `gorm:"column:group_id" json:"group_id"`
	UserID      uint    `gorm:"column:user_id" json:"user_id"`
	Amount      float64 `gorm:"column:amount" json:"amount"`
	Description string  `gorm:"column:description" json:"description"`
	SplitType   string  `gorm:"column:split_type" json:"split_type"` // "equal", "exact", "percentage"
	Status      string  `gorm:"column:status" json:"status"`         // "pending" or "settled"
	Shares      []Share `gorm:"foreignKey:SpendID"`

	//Shares      []Share `gorm:"column:shares" gorm:"foreignKey:SpendID"`
}

type Share struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SpendID   uint      `gorm:"column:spend_id" json:"spend_id"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	Amount    float64   `gorm:"column:amount" json:"amount"`
	Status    string    `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
