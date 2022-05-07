package repository

import (
	"context"
	"github.com/skullkon/info-app/internal/domain"
	"github.com/skullkon/info-app/pkg/client"
	"github.com/skullkon/info-app/pkg/logging"
)

type Information interface {
	Insert(ctx context.Context, info []domain.Info) error
	GetAll(ctx context.Context) ([]domain.Info, error)
	GetRating(ctx context.Context) ([]string, error)
	GetRatingWithParam(ctx context.Context, param string) ([]string, error)
}

type Repositories struct {
	Information Information
}

func NewRepositories(db client.Client, logger *logging.Logger) *Repositories {
	return &Repositories{
		Information: NewRepository(db, logger),
	}
}
