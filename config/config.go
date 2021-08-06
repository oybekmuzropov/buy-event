package config

import (
	"os"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string // develop, staging, production
	LogLevel    string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	//SMTP
	SMTPKEY string

	//Twilio Account information
	AccountSID  string
	AuthToken   string
	TwilioPhone string

	KafkaHost string
	KafkaPort int
}

const (
	MessageTypeSms   = "sms"
	MessageTypeEmail = "email"
)

var MessageTypes = []string{MessageTypeSms, MessageTypeEmail}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))

	c.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", "5432"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "postgres"))
	c.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres"))

	c.SMTPKEY = cast.ToString(getOrReturnDefaultValue("SMTP_KEY", "key"))

	//Twilio
	c.AccountSID = cast.ToString(getOrReturnDefaultValue("ACCOUNT_SID", "sid"))
	c.AuthToken = cast.ToString(getOrReturnDefaultValue("AUTH_TOKEN", "token"))
	c.TwilioPhone = cast.ToString(getOrReturnDefaultValue("TWILIO_PHONE", "phone"))

	c.KafkaHost = cast.ToString(getOrReturnDefaultValue("KAFKA_HOST", "localhost"))
	c.KafkaPort = cast.ToInt(getOrReturnDefaultValue("KAFKA_HOST", 9092))

	return c
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
