package handlers

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/services"
	"strconv"
)

func GetRemainingBalance(c *fiber.Ctx) error {
	groupID, err := strconv.Atoi(c.Params("group_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group ID"})
	}

	balance, err := services.CalculateRemainingBalance(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch balance"})
	}
	return c.JSON(fiber.Map{"balance": balance})
}
