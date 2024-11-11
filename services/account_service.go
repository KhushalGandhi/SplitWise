package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"splitwise/models"
	"splitwise/repositories"
)

func RegisterAccount(account *models.AccountRequest) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	baseModel := models.Account{
		Model:    gorm.Model{},
		ID:       0,
		Name:     account.Name,
		Email:    account.Email,
		Password: account.Password,
	}
	return repositories.CreateAccount(&baseModel)
}

func Authenticate(email, password string) (models.Account, error) {
	account, err := repositories.GetAccountByEmail(email)
	if err != nil {
		return models.Account{}, err
	}
	if bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)) != nil {
		return models.Account{}, errors.New("invalid credentials")
	}
	return account, nil
}
