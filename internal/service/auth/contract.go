//go:generate gomock --source
package auth

import (
	"context"

	"go-atch-template/internal/model"
)

type userStorage interface {
	GetUserByLogin(ctx context.Context, login string) (model.User, error)
}

type notifier interface {
	NotifyUser(user model.User) error
}
