package middleware

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/utils"
	"strings"
)

func JWTAuth(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing or malformed JWT",
		})
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	userID, err := utils.ValidateJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid or expired JWT",
		})
	}

	// Assuming `userID` is a uint or can be converted to one
	c.Locals("userID", uint(userID))
	return c.Next()
}
