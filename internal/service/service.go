package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/skullkon/info-app/internal/domain"
	"github.com/skullkon/info-app/internal/repository"
	"github.com/skullkon/info-app/pkg/logging"
)

type Information interface {
	GetAll(ctx context.Context) ([]domain.Info, error)
	GetRating(ctx context.Context, attr string) ([]string, error)
	GetRatingWithParam(ctx context.Context, column string, value string, attr string) ([]string, error)
	SendData(ctx context.Context, info domain.ClientInfo, ua string) (uuid.UUID, error)
}

type Services struct {
	Information Information
}

type Deps struct {
	Repos  *repository.Repositories
	Logger logging.Logger
}

func NewServices(deps Deps) *Services {
	informationService := NewInformationService(deps.Repos.Information, &deps.Logger)
	return &Services{Information: informationService}
}
