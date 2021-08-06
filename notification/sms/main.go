package sms

import (
	"errors"
	"github.com/sfreiberg/gotwilio"
)

type Twilio struct {
	SID           string
	Token         string
	Phone         string
	ToPhoneNumber string
}

func (s Twilio) Send(text string) error {
	twilio := gotwilio.NewTwilioClient(s.SID, s.Token)

	_, ex, err := twilio.SendSMS(s.Phone, s.ToPhoneNumber, text, "", "")

	if ex != nil {
		return errors.New(ex.Error())
	}

	if err != nil {
		return err
	}

	return nil
}
