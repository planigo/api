package server

import (
	"database/sql"
	"planigo/internal/services"
	"planigo/pkg/mail"
	"planigo/pkg/store"
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
	mailer := mail.New()
	store2 := store.NewStore(db)

	return &Handlers{
		UserHandler:        &services.UserHandler{Store: store2, Mailer: mailer},
		AuthHandler:        &services.AuthHandler{Store: store2, Mailer: mailer},
		ShopHandler:        &services.ShopHandler{Store: store2},
		HourHandler:        &services.HourHandler{Store: store2},
		ServiceHandler:     &services.ServiceHandler{Store: store2},
		CategoryHandler:    &services.CategoryHandler{Store: store2},
		ReservationHandler: &services.ReservationHandler{Store: store2, Mailer: mailer},
	}
}
