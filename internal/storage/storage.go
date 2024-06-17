package storage

import (
	"context"
	"time"

	"go-atch-template/internal/model"
)

type UserStorage interface {
	GetUserByLogin(ctx context.Context, login string) (model.User, error)
	CreateUser(ctx context.Context, id string) error
}

type StatsStorage interface {
	ReadStats(ctx context.Context, from, to time.Time) ([]model.StatsItem, error)
	SaveStats(ctx context.Context, stats []model.StatsItem) error
}
