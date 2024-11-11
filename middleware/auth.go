package middleware

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/utils"
	"strings"
)

func JWTAuth(c *fiber.Ctx) error {
	tokenString := strings.Replace(c.Get("Authorization"), "Bearer ", "", 1)

	claims, err := utils.ParseJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	c.Locals("userID", claims["user_id"])
	return c.Next()
}
