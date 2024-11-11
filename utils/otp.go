package utils

import (
	"math/rand"
	"time"
)

func GenerateOTP() int {
	rand.Seed(time.Now().UnixNano())
	return 100000 + rand.Intn(900000) // generates a 6-digit OTP
}
