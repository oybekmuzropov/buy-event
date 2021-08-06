package messagebroker

type Producer interface {
	Start() error
	Stop() error
	Publish(key string, body []byte) error
}
