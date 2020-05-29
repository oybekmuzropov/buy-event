package repo

import "github.com/google/uuid"

type Purchase struct {
	ID uuid.UUID
	UserID uuid.UUID
	Goods string
	Price float64
}

type PurchaseStorageI interface {
	Create(good *Purchase) error
	Get(id uuid.UUID) (*Purchase, error)
	GetAll() ([]*Purchase, error)
	Delete(id uuid.UUID) error
}