package main

import (
	"cursos-api/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main(){
	
	config.ConnectDB()

	app := fiber.New()

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}