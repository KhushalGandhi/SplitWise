package utils

import (
	"github.com/segmentio/ksuid"
	"math/rand"
	"time"
)

func GenerateOTP() int {
	rand.Seed(time.Now().UnixNano())
	return 100000 + rand.Intn(900000) // generates a 6-digit OTP
}

func GenerateID() string {
	//"github.com/segmentio/ksuid"
	return ksuid.New().String()
}

func GenerateGroupID() string {
	return "Group_" + GenerateID()
}

func GenerateSpendID() string {
	return "Spend_" + GenerateID()
}
