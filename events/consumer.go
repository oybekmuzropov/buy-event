package events

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/buy_event/storage"

	"github.com/segmentio/kafka-go"
	"github.com/buy_event/config"
	"github.com/buy_event/events/handlers"
	loggerPkg "github.com/buy_event/pkg/logger"
	messageBrokerPkg "github.com/buy_event/pkg/messagebroker"
)

type KafkaConsumer struct {
	kafkaReader  *kafka.Reader
	eventHandler *handlers.EventHandler
	logger       loggerPkg.Logger
}

func NewKafkaConsumer(db *sqlx.DB, conf *config.Config, logger loggerPkg.Logger, topic string, publisher map[string]messageBrokerPkg.Producer) messageBrokerPkg.Consumer {
	connString := fmt.Sprintf("%s:%d", conf.KafkaHost, conf.KafkaPort)

	return &KafkaConsumer{
		kafkaReader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{connString},
			GroupID:        "buy_event",
			Topic:          topic,
			MinBytes:       10e3, // 10KB
			MaxBytes:       10e6, // 10MB
			Partition:      0,
			CommitInterval: time.Second,
		}),
		eventHandler: handlers.NewEventHandler(storage.NewStoragePg(db), logger, *conf, publisher),
		logger:       logger,
	}
}

func (k KafkaConsumer) Start() {
	for {
		m, err := k.kafkaReader.ReadMessage(context.Background())

		if err != nil {
			k.logger.Error("Error on consuming a message:", loggerPkg.Error(err))
			continue
		}

		msg, err := k.eventHandler.Handle(m.Topic, m.Value)

		if err != nil {
			fmt.Println()
			k.logger.Error("failed to handle consumed topic:",
				loggerPkg.String("on topic", m.Topic), loggerPkg.Error(err))
			fmt.Println()
		} else {
			fmt.Println()
			k.logger.Info("Successfully consumed message",
				loggerPkg.String("on topic", m.Topic),
				loggerPkg.String("message", msg))
			fmt.Println()
		}
	}
}
