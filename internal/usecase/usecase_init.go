package usecase

import (
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/repository"
	commonconstant "github.com/muhammad-fakhri/go-libs/constant"
	"github.com/muhammad-fakhri/go-libs/log"
)

type usecase struct {
	repo   repository.Repository
	config *usecaseConfig
	logger log.SLogger
}

type usecaseConfig struct {
	activeCountries []commonconstant.Country
}

func NewUsecase(logger log.SLogger, conf *config.Config, repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
		config: &usecaseConfig{
			activeCountries: conf.CountryFilter,
		},
		logger: logger,
	}
}

func InitDependencies(log log.SLogger) Usecase {
	conf := config.Get()
	repository := repository.InitDependencies(log)

	return NewUsecase(log, conf, repository)
}
