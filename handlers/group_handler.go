package handlers

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/models"
	"splitwise/services"
	"strconv"
)

func CreateGroup(c *fiber.Ctx) error {
	var name models.Name
	if err := c.BodyParser(&name); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	baseModel := models.GroupRequest{
		Name:      name.Name,
		CreatorID: c.Locals("userID").(uint),
	}

	//group.CreatorID = c.Locals("userID").(uint)
	if err := services.CreateGroup(&baseModel); err != nil {
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
	var req models.AddUserToGroupRequest

	// Parse and validate request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Retrieve the user ID from the JWT claims (e.g., using middleware or context)
	creatorID := c.Locals("user_id").(string) // assuming user_id was stored in Locals during JWT auth

	// Call service to add user to the group
	err := services.AddUserToGroup(creatorID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User added to group successfully"})
}
