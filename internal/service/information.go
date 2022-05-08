package service

import (
	"context"
	"github.com/skullkon/info-app/internal/domain"
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

func (s *InformationService) GetAll(ctx context.Context) ([]domain.Info, error) {
	test, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return test, nil

}

func (s *InformationService) GetRating(ctx context.Context, attr string) ([]string, error) {
	test, err := s.repository.GetRating(ctx, attr)
	if err != nil {
		return nil, err
	}
	return test, nil

}

func (s *InformationService) GetRatingWithParam(ctx context.Context, column string, value string, attr string) ([]string, error) {
	test, err := s.repository.GetRating(ctx, attr)
	if err != nil {
		return nil, err
	}
	return test, nil

}
