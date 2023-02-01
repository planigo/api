package mail

import (
	"context"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"os"
	"time"
)

type Mailer struct {
	mailgun mailgun.Mailgun
}

type Content struct {
	To      string
	Subject string
	Body    string
}

func New() *Mailer {
	return &Mailer{mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_PRIVATE"))}
}

func (m *Mailer) Send(content Content) {
	msg := m.mailgun.NewMessage(os.Getenv("EMAIL_SENDER"), content.Subject, content.Body, content.To)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := m.mailgun.Send(ctx, msg)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID: %s Resp: %s", id, resp)
}
