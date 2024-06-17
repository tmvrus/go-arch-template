package api

import (
	"context"

	"go-atch-template/internal/model"
)

type auth interface {
	Auth(ctx context.Context, login, password string) (model.AuthState, error)
}
