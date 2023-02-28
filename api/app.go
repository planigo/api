package api

import (
	"fmt"
	"log"
	"os"
	"planigo/api/routes"
	"planigo/config/database"
	"planigo/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
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

	app := fiber.New(fiber.Config{
		AppName: "Planigo",
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")

	handlers := pkg.NewServices(db)

	// Routers
	routes.UserRoutes(api, handlers.UserHandler)
	routes.AuthRoutes(api, handlers.AuthHandler)
	routes.ShopRoutes(api, handlers.ShopHandler)
	routes.HourRoutes(api, handlers.HourHandler)
	routes.ServicesRoutes(api, handlers.ServiceHandler)
	routes.CategoryRoutes(api, handlers.CategoryHandler)
	routes.ReservationRoutes(api, handlers.ReservationHandler)

	// Endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"status": "fail", "message": errorMessage})
	})

	log.Fatal(app.Listen(":" + port))
}
