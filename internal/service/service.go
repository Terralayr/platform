package service

import (
	"github.com/Terralayr/entity-library/entity"
	"github.com/Terralayr/mvp/clients/mvp"
	"github.com/Terralayr/trlyb/clock"
)

type Service struct {
	Mvp *mvp.Client
}

func New() (*Service, error) {
	mvpClient, err := mvp.NewClient(
		"gateway:8000",
		"bob",
		"password",
	)
	if err != nil {
		return nil, err
	}

	return &Service{
		Mvp: mvpClient,
	}, nil
}

func (s *Service) CreateUser() (*entity.UserID, error) {
	return s.Mvp.CreateUser(
		"alice",
		"alice@dummy.com",
		"password",
	)
}

func (s *Service) CreatePhysicalAsset() (*entity.PhysicalAsset, error) {
	return s.Mvp.CreatePhysicalAsset(
		"test-asset",
		1000,
		1000,
		1000,
		"coucou@coucou.com",
	)
}

func (s *Service) CreateBlock(
	physicalAssetID entity.PhysicalAssetID,
	userID entity.UserID,
) (*mvp.UserBlock, error) {
	clk := clock.NewStd()

	start := clk.StartOfTomorrow()
	end := clk.EndOfTomorrow()

	return s.Mvp.CreateBlock(
		physicalAssetID,
		userID,
		start,
		end,
		500,
		500,
		500,
	)
}
