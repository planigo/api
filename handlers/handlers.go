package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2/middleware/session"
	"planigo/config/mail"
	storeManager "planigo/config/store"
	"time"
)

type Handlers struct {
	UserHandler    *UserHandler
	AuthHandler    *AuthHandler
	ShopHandler    *ShopHandler
	HourHandler    *HourHandler
	ServiceHandler *ServiceHandler
}

func New(db *sql.DB) *Handlers {
	sessionConfig := session.Config{Expiration: 48 * time.Hour}
	mailer := mail.New()
	store := storeManager.NewStore(db)
	s := session.New(sessionConfig)

	return &Handlers{
		UserHandler:    &UserHandler{Store: store, Mailer: mailer, Session: s},
		AuthHandler:    &AuthHandler{Store: store, Mailer: mailer, Session: s},
		ShopHandler:    &ShopHandler{Store: store, Session: s},
		HourHandler:    &HourHandler{Store: store, Session: s},
		ServiceHandler: &ServiceHandler{Store: store, Session: s},
	}
}
