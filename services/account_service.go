package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"splitwise/models"
	"splitwise/repositories"
	"time"
)

func RegisterAccount(account *models.AccountRequest) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	baseModel := models.Account{
		Name:      account.Name,
		Email:     account.Email,
		Password:  account.Password,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
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
