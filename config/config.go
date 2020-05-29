package config

import (
	"os"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	//SMTP
	SMTPKEY		string

	//Twilio Account information
	AccountSID 		string
	AuthToken		string
	TwilioPhone		string
}

const MESSAGE_TYPE_SMS = "sms"
const MESSAGE_TYPE_EMAIL = "email"
var MessageTypes = []string{MESSAGE_TYPE_SMS, MESSAGE_TYPE_EMAIL}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", "5432"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "event"))
	c.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "event"))
	c.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "event123"))

	c.SMTPKEY = cast.ToString(getOrReturnDefaultValue("SMTP_KEY", "key"))

	//Twilio
	c.AccountSID = cast.ToString(getOrReturnDefaultValue("ACCOUNT_SID", "sid"))
	c.AuthToken = cast.ToString(getOrReturnDefaultValue("AUTH_TOKEN", "token"))
	c.TwilioPhone = cast.ToString(getOrReturnDefaultValue("TWILIO_PHONE", "phone"))


	return c
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}