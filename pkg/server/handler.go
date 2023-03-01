package server

import (
	"database/sql"
	"github.com/gofiber/fiber/v2/middleware/session"
	"planigo/internal/services"
	"planigo/pkg/mail"
	"planigo/pkg/store"
	"time"
)

type Handlers struct {
	UserHandler        *services.UserHandler
	AuthHandler        *services.AuthHandler
	ShopHandler        *services.ShopHandler
	HourHandler        *services.HourHandler
	ServiceHandler     *services.ServiceHandler
	CategoryHandler    *services.CategoryHandler
	ReservationHandler *services.ReservationHandler
}

func RegisterHandlers(db *sql.DB) *Handlers {
	sessionConfig := session.Config{
		Expiration: 48 * time.Hour,
	}
	mailer := mail.New()
	store2 := store.NewStore(db)
	sess := session.New(sessionConfig)

	return &Handlers{
		UserHandler:        &services.UserHandler{Store: store2, Mailer: mailer, Session: sess},
		AuthHandler:        &services.AuthHandler{Store: store2, Mailer: mailer, Session: sess},
		ShopHandler:        &services.ShopHandler{Store: store2, Session: sess},
		HourHandler:        &services.HourHandler{Store: store2, Session: sess},
		ServiceHandler:     &services.ServiceHandler{Store: store2, Session: sess},
		CategoryHandler:    &services.CategoryHandler{Store: store2},
		ReservationHandler: &services.ReservationHandler{Store: store2, Session: sess, Mailer: mailer},
	}
}
