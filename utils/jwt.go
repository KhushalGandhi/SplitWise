package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

//func ParseJWT(tokenString string) (jwt.MapClaims, error) {
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte(os.Getenv("JWT_SECRET")), nil
//	})
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		return claims, nil
//	}
//	return nil, err
//}

func ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, ok := claims["user_id"].(float64) // JWT stores numbers as float64
		if !ok {
			return 0, errors.New("invalid token claims")
		}
		return uint(userID), nil
	}

	return 0, errors.New("invalid token")
}
