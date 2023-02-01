package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"planigo/api/routes"
	"planigo/config/database"
	store2 "planigo/config/store"
	"planigo/pkg/user"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := fiber.New()
	app.Use(logger.New())
	api := app.Group("/api")

	store := store2.NewStore(db)
	handler := &user.Handler{Store: store}

	routes.UserRoutes(api, handler)

	log.Fatal(app.Listen(":8080"))
}
