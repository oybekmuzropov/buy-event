package storage

import (
	"github.com/buy_event/storage/potgres"
	"github.com/buy_event/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
	Purchase() repo.PurchaseStorageI
	Log() repo.LogStorageI
}

type storagePg struct {
	db *sqlx.DB
	userRepo repo.UserStorageI
	purchaseRepo repo.PurchaseStorageI
	logRepo repo.LogStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:db,
		userRepo:potgres.NewUserRepo(db),
		purchaseRepo:potgres.NewPurchaseRepo(db),
		logRepo:potgres.NewLogRepo(db),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s storagePg) Purchase() repo.PurchaseStorageI {
	return s.purchaseRepo
}

func (s storagePg) Log() repo.LogStorageI {
	return s.logRepo
}