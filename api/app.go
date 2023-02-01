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

	api := app.Group("/api")

	routes.UserRoutes(api)
	//routes.UserRoutes(router)
	//routes.UserRoutes(router)
	//routes.UserRoutes(router)

	log.Fatal(app.Listen(":8080"))
}
