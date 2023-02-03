package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"planigo/api/routes"
	"planigo/config/database"
	"planigo/handlers"
)

func Start() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := fiber.New(fiber.Config{
		AppName: "Planigo",
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")

	h := handlers.New(db)

	// Routers
	routes.UserRoutes(api, h.UserHandler)
	routes.AuthRoutes(api, h.AuthHandler)
	routes.ShopRoutes(api, h.ShopHandler)
	routes.HourRoutes(api, h.HourHandler)
	routes.ServicesRoutes(api, h.ServiceHandler)

	// Endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"status": "fail", "message": errorMessage})
	})

	log.Fatal(app.Listen(":8080"))
}
