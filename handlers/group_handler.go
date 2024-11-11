package handlers

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/models"
	"splitwise/services"
)

func CreateGroup(c *fiber.Ctx) error {
	var group models.GroupRequest
	if err := c.BodyParser(&group); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	group.CreatorID = c.Locals("userID").(uint)
	if err := services.CreateGroup(&group); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create group"})
	}
	return c.JSON(fiber.Map{"message": "Group created successfully"})
}
