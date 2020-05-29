package cmd

import (
	"errors"
	"github.com/sfreiberg/gotwilio"
)

func sendSMS(toPhoneNumber, text string) error {
	twilio := gotwilio.NewTwilioClient(cfg.AccountSID, cfg.AuthToken)

	_, ex, err := twilio.SendSMS(cfg.TwilioPhone, toPhoneNumber, text, "", "")

	if ex != nil {
		return errors.New(ex.Error())
	}

	return err
}
