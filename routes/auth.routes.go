package routes

import (
	auth "cursos-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/auth")

	// Route for user registration
	api.Post("/signup", auth.SingUp)
}