package handlers

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/models"
	"splitwise/services"
	"strconv"
)

func CreateSpend(c *fiber.Ctx) error {

	groupId := c.Params("id")
	userId := c.Locals("userID").(uint)

	var spend models.CreateSpendRequest

	spend.GroupID = groupId
	spend.UserID = userId
	if err := c.BodyParser(&spend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := services.CreateSpend(spend); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create spend"})
	}
	return c.JSON(fiber.Map{"message": "Spend added successfully"})
}

func MarkSharePaidHandler(c *fiber.Ctx) error {
	spendID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid spend ID"})
	}

	userID := c.Locals("userID").(uint)

	if err := services.MarkShareAsPaid(uint(spendID), userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "share marked as paid and spend updated"})
}
