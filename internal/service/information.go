package service

import (
	"context"
	"github.com/skullkon/info-app/internal/repository"
	"github.com/skullkon/info-app/pkg/logging"
)

type InformationService struct {
	repository repository.Information
	logger     *logging.Logger
}

func NewInformationService(repo repository.Information, logger *logging.Logger) *InformationService {
	return &InformationService{
		repository: repo,
		logger:     logger,
	}
}

func (s *InformationService) HelloWorld(ctx context.Context) (string, error) {
	return "hello world", nil
}
