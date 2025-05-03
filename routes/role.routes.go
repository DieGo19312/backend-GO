package routes

import (
	role "cursos-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoleRoutes(app *fiber.App) {
	api := app.Group("/role")
	api.Post("/add", role.CreateRole)
}