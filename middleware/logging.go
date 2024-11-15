package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"splitwise/models"
	"splitwise/repositories"
)

func Logging(c *fiber.Ctx) error {
	err := c.Next()

	logEntry := models.Log{
		Endpoint: c.Path(),
		Method:   c.Method(),
		Status:   c.Response().StatusCode(),
	}

	if err := repositories.CreateLog(&logEntry); err != nil {
		log.Println("Failed to log request:", err)
	}

	return err
}
