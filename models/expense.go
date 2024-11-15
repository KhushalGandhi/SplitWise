package models

type CompletePaymentRequest struct {
	SpendID uint `json:"spend_id"`
	UserID  uint `json:"user_id"`
}
