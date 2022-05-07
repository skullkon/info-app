package service

import (
	"context"
	"github.com/skullkon/info-app/internal/repository"
)

type Information interface {
	HelloWorld(ctx context.Context) (string, error)
}

type Services struct {
	Information Information
}

type Deps struct {
	Repos *repository.Repositories
}
