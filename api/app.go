package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	"log"
	"planigo/api/routes"
	"planigo/config/database"
	"planigo/config/mail"
	storeManager "planigo/config/store"
	"planigo/pkg/auth"
	"planigo/pkg/user"
	"time"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := fiber.New()

	sessionConfig := session.Config{Expiration: 48 * time.Hour}

	store := storeManager.NewStore(db)
	mailer := mail.New()
	session2 := session.New(sessionConfig)

	// Middlewares
	app.Use(logger.New())

	api := app.Group("/api")

	// Controllers
	userHandler := &user.Handler{Store: store, Mailer: mailer, Session: session2}
	authHandler := &auth.Handler{Store: store, Mailer: mailer, Session: session2}

	// Routers
	routes.UserRoutes(api, userHandler)
	routes.AuthRoutes(api, authHandler)

	log.Fatal(app.Listen(":8080"))
}
