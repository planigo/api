package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"planigo/api/routes"
	"planigo/config/database"
	"planigo/config/mail"
	store2 "planigo/config/store"
	"planigo/pkg/user"
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
	app.Use(logger.New())
	api := app.Group("/api")

	store := store2.NewStore(db)
	mailer := mail.New()
	handler := &user.Handler{Store: store, Mailer: mailer}

	routes.UserRoutes(api, handler)

	log.Fatal(app.Listen(":8080"))
}
