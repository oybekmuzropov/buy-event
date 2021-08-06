package handlers

import (
	"encoding/json"
	"fmt"
	configPkg "github.com/buy_event/config"
	"github.com/buy_event/entities"
	"github.com/buy_event/notification"
	loggerPkg "github.com/buy_event/pkg/logger"
	"github.com/buy_event/pkg/messagebroker"
	"github.com/buy_event/storage"
	"github.com/buy_event/storage/repo"
)

type EventHandler struct {
	conf      *configPkg.Config
	storage   storage.StorageI
	logger    loggerPkg.Logger
	publisher map[string]messagebroker.Producer
}

func NewEventHandler(storage storage.StorageI, logger loggerPkg.Logger, conf configPkg.Config, publisher map[string]messagebroker.Producer) *EventHandler {
	return &EventHandler{
		storage:   storage,
		conf:      &conf,
		logger:    logger,
		publisher: publisher,
	}
}

func (h *EventHandler) Handle(topic string, value []byte) (string, error) {
	switch topic {
	case "notification":
		var n entities.Notification

		err := json.Unmarshal(value, &n)
		if err != nil {
			h.logger.Error("failed to unmarshal byte to notification", loggerPkg.Error(err))
			return "", err
		}

		notification, err := notification.NewNotification(
			topic,
			notification.Credential{
				EmailCredential: notification.EmailCredential{
					SMTPKey: h.conf.SMTPKEY,
					ToEmail: n.Email,
				},
				SMSCredential: notification.SMSCredential{
					AccountSID:    h.conf.AccountSID,
					AuthToken:     h.conf.AuthToken,
					TwilioPhone:   h.conf.TwilioPhone,
					ToPhoneNumber: n.PhoneNumber,
				},
			})

		if err != nil {
			h.logger.Error("failed to create instance of notification", loggerPkg.Error(err))
			return "", err
		}

		err = notification.Send(n.Text)
		if err != nil {
			log := entities.Log{
				Error:      err.Error(),
				PurchaseID: n.PurchaseID,
			}

			body, err := json.Marshal(log)
			if err != nil {
				h.logger.Error("failed marshal log to []byte", loggerPkg.Error(err))
				return "", err
			}

			err = h.publisher["log"].Publish("log", body)
			if err != nil {
				return "", err
			}
		}
	case "log":
		var log entities.Log

		err := json.Unmarshal(value, &log)
		if err != nil {
			h.logger.Error("failed to unmarshal byte to log", loggerPkg.Error(err))
			return "", err
		}

		err = h.storage.Log().Create(&repo.Log{
			Error:      log.Error,
			PurchaseID: log.PurchaseID,
		})
		if err != nil {
			h.logger.Error("failed while creating log", loggerPkg.Error(err))
		}
	}

	return fmt.Sprintf("message consumed successfully: %s", topic), nil
}
