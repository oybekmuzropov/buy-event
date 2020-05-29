package repo

import "github.com/google/uuid"

type Log struct {
	ID uuid.UUID
	Error string
}

type LogStorageI interface {
	Create(good *Log) error
	GetAll() ([]*Log, error)
}
