package handlers

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/models"
	"splitwise/services"
)

func CreateSpend(c *fiber.Ctx) error {
	var spend *models.CreateSpendRequest
	if err := c.BodyParser(&spend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := services.CreateSpend(spend); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create spend"})
	}
	return c.JSON(fiber.Map{"message": "Spend added successfully"})
}
