package stats

import (
	"context"
	"time"

	"go-atch-template/internal/model"
)

type storage struct{}

func (s *storage) ReadStats(ctx context.Context, from, to time.Time) ([]model.StatsItem, error) {
	return nil, nil
}

func (s *storage) SaveStats(ctx context.Context, stats []model.StatsItem) error {
	return nil
}
