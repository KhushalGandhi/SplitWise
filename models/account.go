package models

import "time"

type Account struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"name" json:"name"`
	Email     string    `gorm:"email;unique" json:"email"`
	Password  string    `gorm:"password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type AccountRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
