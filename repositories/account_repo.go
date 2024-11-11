package repositories

import (
	"splitwise/db"
	"splitwise/models"
)

func CreateAccount(account *models.Account) error {
	return database.DB.Create(account).Error
}

func GetAccountByEmail(email string) (models.Account, error) {
	var account models.Account
	err := database.DB.Where("email = ?", email).First(&account).Error
	return account, err
}
