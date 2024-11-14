package handlers

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/services"
)

func GetRemainingBalance(c *fiber.Ctx) error {
	groupID := c.Params("group_id")
	//userId := c.Locals("userId").(uint)

	balance, err := services.CalculateRemainingBalance(groupID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch balance"})
	}
	return c.JSON(fiber.Map{"balance": balance})
}
