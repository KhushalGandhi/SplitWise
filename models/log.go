package models

import "time"

type Log struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Endpoint  string    `json:"endpoint"`
	Method    string    `json:"method"`
	Status    int       `json:"status"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
