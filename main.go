package main

import (
	"cursos-api/config"
	"cursos-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main(){

	config.ConnectDB()

	app := fiber.New()

	routes.AuthRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}