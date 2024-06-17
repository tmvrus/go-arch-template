package auth

import (
	"context"
	"fmt"

	"go-atch-template/internal/model"
	"go-atch-template/internal/storage"
)

type service struct {
	storage userStorage
	n       notifier
}

func New(s storage.UserStorage, n notifier) service {
	return service{storage: s, n: n}
}

func (s service) Auth(ctx context.Context, login, password string) (model.AuthState, error) {
	user, err := s.storage.GetUserByLogin(ctx, login)
	if err != nil {
		return model.AuthState{}, err
	}

	if user.Password != password {
		return model.AuthState{}, fmt.Errorf("invalid password")
	}

	state := model.AuthState{Valid: true}
	err = s.n.NotifyUser(user)
	if err != nil {
		// log err
	}

	return state, nil
}
