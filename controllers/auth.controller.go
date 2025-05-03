package controllers

import (
	"cursos-api/config"
	"cursos-api/models"
	"cursos-api/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SingUp (c *fiber.Ctx) error {
	db := config.GetDB()
	if db == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database connection is not established")
	}

	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Error parsing request body")
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Name, email and password are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error hashing password")
	}
	user.Password = string(hashedPassword)
	user.Active = true
	
	if user.RoleID == nil {
		defaultRoleID := 1
		user.RoleID = &defaultRoleID
	}
	user.ID = utils.RandomString(16)

	data := map[string]interface{}{
		"id":          user.ID,
		"name":        user.Name,
		"pat_summary": user.PatSummary,
		"mat_summary": user.MatSummary,
		"username":    user.Username,
		"age":         user.Age,
		"email":       user.Email,
		"password":    user.Password,
		"phone":       user.Phone,
		"role_id":     *user.RoleID,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"active":     user.Active,
}
	query, values := utils.ParseInsertArray("users", data)

	err = db.QueryRow(query+" RETURNING id, created_at", values...).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "User already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"id":         user.ID,
			"name":       user.Name,
			"pat_summary": user.PatSummary,
			"mat_summary": user.MatSummary,
			"username":   user.Username,
			"age":        user.Age,
			"email":      user.Email,
			"phone":      user.Phone,
			"created_at": user.CreatedAt,
			},
		})
}