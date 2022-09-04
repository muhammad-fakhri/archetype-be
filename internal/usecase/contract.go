package usecase

import (
	"context"

	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
)

type Usecase interface {
	// BEGIN __INCLUDE_EXAMPLE_CRON__
	UpdateEventExample(ctx context.Context) (err error)
	// END __INCLUDE_EXAMPLE_CRON__
	// BEGIN __INCLUDE_EXAMPLE__
	UpdateSystemConfig(ctx context.Context, configs *dto.PublicSystemConfigRequest) (err error)
	GetSystemConfig(ctx context.Context) (configs *dto.PublicSystemConfig, err error)
	UpdateSystemConfigs(ctx context.Context, configMap map[constant.SystemConfig]interface{}) (err error)
	// END __INCLUDE_EXAMPLE__
}
