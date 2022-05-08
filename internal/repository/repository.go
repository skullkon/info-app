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
	GetRating(ctx context.Context, attr string) ([]string, error)
	GetRatingWithParam(ctx context.Context, column string, value string, attr string) ([]string, error)
	IpExist(ctx context.Context, ip string) (bool, error)
	GetIdByIp(ctx context.Context, ip string) (int32, error)
}

type Repositories struct {
	Information Information
}

func NewRepositories(db client.Client, logger *logging.Logger) *Repositories {
	return &Repositories{
		Information: NewRepository(db, logger),
	}
}
