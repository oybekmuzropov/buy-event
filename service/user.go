package service

import (
	"context"
	"github.com/buy_event/storage"
	"github.com/buy_event/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserService struct {
	storage storage.StorageI
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
	}
}

func (us *UserService) Create(ctx context.Context, req *repo.User) error {
	err := us.storage.User().Create(req)

	return err
}

func (us *UserService) Get(ctx context.Context, id uuid.UUID) (*repo.User, error) {
	user, err := us.storage.User().Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetAll(ctx context.Context) ([]*repo.User, error) {
	users, err := us.storage.User().GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	err := us.storage.User().Delete(id)

	return err
}