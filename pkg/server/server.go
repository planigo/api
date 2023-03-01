package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"planigo/config/database"
	handlers2 "planigo/internal/handlers"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := fiber.New(fiber.Config{AppName: "Planigo"})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")

	handlers := RegisterHandlers(db)

	// Routers
	handlers2.UserRoutes(api, handlers.UserHandler)
	handlers2.AuthRoutes(api, handlers.AuthHandler)
	handlers2.ShopRoutes(api, handlers.ShopHandler)
	handlers2.HourRoutes(api, handlers.HourHandler)
	handlers2.ServicesRoutes(api, handlers.ServiceHandler)
	handlers2.CategoryRoutes(api, handlers.CategoryHandler)
	handlers2.ReservationRoutes(api, handlers.ReservationHandler)

	// Endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"status": "fail", "message": errorMessage})
	})

	log.Fatal(app.Listen(":" + port))
}
