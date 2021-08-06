package notification

import (
	"errors"
	"github.com/buy_event/config"
	"github.com/buy_event/notification/email"
	"github.com/buy_event/notification/sms"
)

type INotification interface {
	Send(text string) error
}

type EmailCredential struct {
	SMTPKey string
	ToEmail string
}

type SMSCredential struct {
	AccountSID    string
	AuthToken     string
	TwilioPhone   string
	ToPhoneNumber string
}

type Credential struct {
	EmailCredential EmailCredential
	SMSCredential   SMSCredential
}

func NewNotification(notificationType string, credential Credential) (INotification, error) {
	switch notificationType {
	case config.MessageTypeEmail:
		return email.Sendgrid{
			Key: credential.EmailCredential.SMTPKey,
		}, nil
	case config.MessageTypeSms:
		return sms.Twilio{
			SID:   credential.SMSCredential.AccountSID,
			Token: credential.SMSCredential.AuthToken,
			Phone: credential.SMSCredential.TwilioPhone,
		}, nil
	}

	return nil, errors.New("sms_type not found")
}
