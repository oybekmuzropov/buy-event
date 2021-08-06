package service

import (
	"context"
	"github.com/buy_event/storage"
	"github.com/buy_event/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PurchaseService struct {
	storage storage.StorageI
}

func NewPurchaseService(db *sqlx.DB) *PurchaseService {
	return &PurchaseService{storage: storage.NewStoragePg(db)}
}

func (ps *PurchaseService) Create(ctx context.Context, req *repo.Purchase) (string, error) {
	guid, err := ps.storage.Purchase().Create(req)

	if err != nil {
		return "", err
	}

	return guid, nil
}

func (ps *PurchaseService) Get(ctx context.Context, id uuid.UUID) (*repo.Purchase, error) {
	purchase, err := ps.storage.Purchase().Get(id)

	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (ps *PurchaseService) GetAll(ctx context.Context) ([]*repo.Purchase, error) {
	purchases, err := ps.storage.Purchase().GetAll()

	if err != nil {
		return nil, err
	}

	return purchases, nil
}

func (ps *PurchaseService) Delete(ctx context.Context, id uuid.UUID) error {
	err := ps.storage.Purchase().Delete(id)

	return err
}
