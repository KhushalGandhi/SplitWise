package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID       uint   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type AccountRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
