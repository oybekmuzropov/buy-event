package email

import (
	"errors"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
)

type Sendgrid struct {
	Key     string
	ToEmail string
}

func (e Sendgrid) Send(text string) error {
	from := mail.NewEmail("Oybek Muzropov", "oybekmuzropov@yandex.com")
	subject := "Your purchase report"
	to := mail.NewEmail("User", e.ToEmail)
	plainTextContent := text
	htmlContent := fmt.Sprintf("<strong>%s</strong>", text)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(e.Key)
	res, err := client.Send(message)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusAccepted {
		return errors.New(res.Body)
	}

	return nil
}
