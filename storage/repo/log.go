package repo

import "github.com/google/uuid"

type Log struct {
	ID        uuid.UUID
	Error     string
	PurchaseID string
}

type LogStorageI interface {
	Create(log *Log) error
	GetAll() ([]*Log, error)
}
