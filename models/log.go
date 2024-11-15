package models

import "time"

type Log struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Endpoint  string    `gorm:"endpoint" json:"endpoint"`
	Method    string    `gorm:"method" json:"method"`
	Status    int       `gorm:"status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
