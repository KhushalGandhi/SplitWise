package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
	Status   int    `json:"status"`
	UserID   uint   `json:"user_id"`
}
