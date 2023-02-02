package api

import (
	"fmt"
	"log"
	"planigo/api/routes"
	"planigo/config/database"
	"planigo/config/mail"
	storeManager "planigo/config/store"
	"planigo/pkg/auth"
	"planigo/pkg/shop"
	"planigo/pkg/user"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
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

	sessionConfig := session.Config{Expiration: 48 * time.Hour}

	store := storeManager.NewStore(db)
	mailer := mail.New()
	session := session.New(sessionConfig)

	// Middlewares
	app.Use(logger.New())

	api := app.Group("/api")

	// Controllers
	userHandler := &user.Handler{Store: store, Mailer: mailer, Session: session}
	authHandler := &auth.Handler{Store: store, Mailer: mailer, Session: session}
	shopHandler := &shop.Handler{Store: store, Session: session}

	// Routers
	routes.UserRoutes(api, userHandler)
	routes.AuthRoutes(api, authHandler)
	routes.ShopRoutes(api, shopHandler)

	// Endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"status": "fail", "message": errorMessage})
	})

	log.Fatal(app.Listen(":8080"))
}
