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

func GetAccountIdByEmail(email string) (uint, error) {
	var account models.Account
	err := database.DB.Where("email = ?", email).First(&account).Error
	return account.ID, err
}

func GetAccountByID(userID uint) (*models.Account, error) {
	var account models.Account
	if err := database.DB.Where("id = ?", userID).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func GetAccountByNameAndEmail(email string, name string) (*models.Account, error) {
	var account models.Account
	if err := database.DB.Where("email ILIKE ? AND name ILIKE ?", email, name).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil

}

func GetEmailByUserID(userID uint) (string, error) {
	var account models.Account
	if err := database.DB.Where("id = ?", userID).First(&account).Error; err != nil {
		return "", err
	}
	return account.Email, nil
}
