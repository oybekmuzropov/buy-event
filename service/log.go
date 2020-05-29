package service

import (
	"context"
	storage "github.com/buy_event/storage"
	"github.com/buy_event/storage/repo"
	"github.com/jmoiron/sqlx"
)

type LogService struct {
	storage storage.StorageI
}

func NewLogService(db *sqlx.DB) *LogService {
	return &LogService{storage:storage.NewStoragePg(db)}
}

func (ls *LogService) Create(ctx context.Context, req *repo.Log) error {
	err := ls.storage.Log().Create(req)

	return err
}

func (ls *LogService) GetAll(ctx context.Context) ([]*repo.Log, error) {
	logs, err := ls.storage.Log().GetAll()

	return logs, err
}
