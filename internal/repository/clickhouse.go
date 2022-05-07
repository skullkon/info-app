package repository

import (
	"context"
	"github.com/skullkon/info-app/internal/domain"
	"github.com/skullkon/info-app/pkg/client"
	"github.com/skullkon/info-app/pkg/logging"
)

type Repository struct {
	client client.Client
	logger *logging.Logger
}

func NewRepository(client client.Client, logger *logging.Logger) *Repository {
	return &Repository{
		client: client,
		logger: logger,
	}
}

func (r *Repository) Insert(ctx context.Context, info []domain.Info) error {

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]domain.Info, error) {
	var test []domain.Info
	err := r.client.Select(ctx, &test, "select * from info limit 10")
	if err != nil {
		r.logger.Error(err)
		return nil, err
	}
	return test, nil
}

func (r *Repository) GetRating(ctx context.Context) ([]string, error) {
	return nil, nil
}

func (r *Repository) GetRatingWithParam(ctx context.Context, param string) ([]string, error) {
	return nil, nil
}
