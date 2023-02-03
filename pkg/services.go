package pkg

import (
	"database/sql"
	"github.com/gofiber/fiber/v2/middleware/session"
	"planigo/config/mail"
	storeManager "planigo/config/store"
	"planigo/pkg/auth"
	"planigo/pkg/hour"
	"planigo/pkg/service"
	"planigo/pkg/reservation"
	"planigo/pkg/shop"
	"planigo/pkg/user"
	"time"
)

type Services struct {
	UserHandler        *user.Handler
	AuthHandler        *auth.Handler
	ShopHandler        *shop.Handler
	ReservationHandler *reservation.Handler
	ServiceHandler *service.ServiceHandler
	HourHandler *hour.Handler
}

func NewServices(db *sql.DB) *Services {
	sessionConfig := session.Config{Expiration: 48 * time.Hour}
	mailer := mail.New()
	store := storeManager.NewStore(db)
	session := session.New(sessionConfig)

	return &Services{
		UserHandler: &user.Handler{Store: store, Mailer: mailer, Session: session},
		AuthHandler: &auth.Handler{Store: store, Mailer: mailer, Session: session},
		ShopHandler: &shop.Handler{Store: store, Session: session},
		HourHandler: &hour.Handler{Store: store, Session: session},
		ServiceHandler: &service.ServiceHandler{Store: store, Session: session},
		ReservationHandler: &reservation.Handler{Store: store, Session: session},
	}
}
