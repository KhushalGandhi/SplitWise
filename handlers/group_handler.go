package handlers

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/models"
	"splitwise/services"
	"strconv"
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

func DeleteGroup(c *fiber.Ctx) error {
	groupID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group ID"})
	}

	userID := c.Locals("userID").(uint)

	if err := services.CanDeleteGroup(uint(groupID), userID); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.DeleteGroup(uint(groupID), userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete group"})
	}

	return c.JSON(fiber.Map{"message": "Group deleted successfully"})
}

func AddUserToGroup(c *fiber.Ctx) error {
	groupID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group ID"})
	}

	userID := c.Locals("userID").(uint)

	if err := services.CanAddUserToGroup(uint(groupID), userID); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	var newUser struct {
		UserID uint `json:"user_id"`
	}
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := services.AddUserToGroup(uint(groupID), newUser.UserID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not add user"})
	}

	return c.JSON(fiber.Map{"message": "User added to group successfully"})
}
