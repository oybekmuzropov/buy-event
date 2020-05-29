package repo

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID
	Email string `binding:"required"`
	PhoneNumber string `binding:"required"`
}

type UserStorageI interface {
	Create(*User) error
	Get(id uuid.UUID) (*User, error)
	GetAll() ([]*User, error)
	Delete(id uuid.UUID) error
}
