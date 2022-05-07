package service

import (
	"context"
	"github.com/skullkon/info-app/internal/repository"
	"github.com/skullkon/info-app/pkg/logging"
)

type Information interface {
	HelloWorld(ctx context.Context) (string, error)
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
