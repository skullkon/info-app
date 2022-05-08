package service

import (
	"context"
	"github.com/skullkon/info-app/internal/domain"
	"github.com/skullkon/info-app/internal/repository"
	"github.com/skullkon/info-app/pkg/logging"
	"github.com/skullkon/info-app/pkg/utils"
	"github.com/ua-parser/uap-go/uaparser"
	"math/rand"
	"time"
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

func (s *InformationService) SendData(ctx context.Context, info domain.ClientInfo, ua string) (int32, error) {
	parser, err := uaparser.New("./regexes.yaml")
	if err != nil {
		s.logger.Fatal(err)
	}
	agent := parser.Parse(ua)

	newInfo := domain.Info{
		Ip:             info.Ip,
		TypeOfDevice:   agent.Device.Family,
		Os:             agent.Os.Family,
		OsVersion:      agent.Os.Major,
		Browser:        agent.UserAgent.Family,
		BrowserVersion: agent.UserAgent.Major + "." + agent.UserAgent.Minor,
		Brand:          agent.Device.Brand,
		Model:          agent.Device.Model,
		Resolution:     utils.GenResolution(),
		Time:           time.Now(),
	}

	ipExist, err := s.repository.IpExist(ctx, info.Ip)
	if err != nil {
		return -1, err
	}

	if info.Id < 0 && ipExist == false {
		// in the whole system I need to replace int id with guid
		//in order to simply generate random user ids, but
		//I forgot about this at the design stage
		rand.Seed(time.Now().UnixNano())
		newInfo.Id = rand.Int31n(10000)

		infoList := []domain.Info{newInfo}
		if err := s.repository.Insert(ctx, infoList); err != nil {
			return -1, err
		}

		return newInfo.Id, nil
	}

	if info.Id < 0 && ipExist == true {
		newInfo.Id, err = s.repository.GetIdByIp(ctx, info.Ip)
		if err != nil {
			return -1, err
		}
		infoList := []domain.Info{newInfo}
		if err := s.repository.Insert(ctx, infoList); err != nil {
			return -1, err
		}

		return newInfo.Id, nil
	}

	return -1, nil
}
