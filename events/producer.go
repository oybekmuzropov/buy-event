package events

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/buy_event/config"
	loggerPkg "github.com/buy_event/pkg/logger"
	brokerPkg "github.com/buy_event/pkg/messagebroker"
	"time"
)

type KafkaProducer struct {
	kafkaWriter *kafka.Writer
	logger      loggerPkg.Logger
}

func NewKafkaProducer(conf config.Config, logger loggerPkg.Logger, topic string) brokerPkg.Producer {
	connString := fmt.Sprintf("%s:%d", conf.KafkaHost, conf.KafkaPort)

	return &KafkaProducer{
		kafkaWriter: kafka.NewWriter(kafka.WriterConfig{
			Brokers:      []string{connString},
			Topic:        topic,
			BatchTimeout: 10 * time.Millisecond,
		}),
		logger: logger,
	}
}

// Start ...
func (p *KafkaProducer) Start() error {
	return nil
}

// Stop ...
func (p *KafkaProducer) Stop() error {
	err := p.kafkaWriter.Close()
	if err != nil {
		return err
	}

	return nil
}

// Publish ...
func (p *KafkaProducer) Publish(key string, body []byte) error {
	message := kafka.Message{
		Key:   []byte(key),
		Value: body,
	}

	if err := p.kafkaWriter.WriteMessages(context.Background(), message); err != nil {
		return err
	}

	return nil
}
