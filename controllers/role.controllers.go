package controllers

import (
	"cursos-api/config"
	"cursos-api/models"
	"cursos-api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	db := config.GetDB()
	if db == nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Database connection error",
		})
	}

	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	data := map[string]interface{}{
		"name":        role.Name,
		"description": role.Description,
		"permissions": role.Permissions,
	}

	query, values := utils.ParseInsertArray("roles", data)
	err := db.QueryRow(query+" RETURNING id", values...).Scan(&role.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error creating role",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Role created successfully",
		"role":    role,
	})
}