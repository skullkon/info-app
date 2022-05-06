package storage

import (
	"context"
	"github.com/skullkon/info-app/internal/information"
	"github.com/skullkon/info-app/pkg/client"
	"github.com/skullkon/info-app/pkg/logging"
)

type repository struct {
	client client.Client
	logger *logging.Logger
}

func NewRepository(client client.Client, logger *logging.Logger) information.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Insert(ctx context.Context, info []information.Info) error {
	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]information.Info, error) {
	return nil, nil
}

func (r *repository) GetRating(ctx context.Context) ([]string, error) {
	return nil, nil
}

func (r *repository) GetRatingWithParam(ctx context.Context, param string) ([]string, error) {
	return nil, nil
}
