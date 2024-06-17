package user

import (
	"context"

	"go-atch-template/internal/model"
)

type storage struct{}

func New(dbAddress string) storage {
	return storage{}
}

func (s storage) GetUserByLogin(ctx context.Context, login string) (model.User, error) {
	return model.User{}, nil
}

func (s storage) CreateUser(ctx context.Context, id string) error {
	return nil
}
