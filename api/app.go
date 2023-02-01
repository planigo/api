package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"planigo/api/routes"
	"planigo/config/database"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())

	router := app.Group("/api")

	router.Get("/users", routes.GetUsers)

	log.Fatal(app.Listen(":8080"))
}
