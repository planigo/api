package api

import (
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

	app := fiber.New()

	session := session.New(session.Config{
		Expiration: 48 * time.Hour,
	})

	mailer := mail.New()
	store := storeManager.NewStore(db)

	app.Use(logger.New())

	// Api routes
	api := app.Group("/api")

	// User routes
	userHandler := &user.Handler{Store: store, Mailer: mailer, Session: session}
	authHandler := &auth.Handler{Store: store, Mailer: mailer, Session: session}
	shopHandler := &shop.ShopHandler{Store: store}

	routes.UserRoutes(api, userHandler)
	routes.AuthRoutes(api, authHandler)
	routes.ShopRoutes(api, shopHandler)

	log.Fatal(app.Listen(":8080"))
}
