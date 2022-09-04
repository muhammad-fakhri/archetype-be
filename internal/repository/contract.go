package repository

import (
	"context"

	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/model"
)

type Repository interface {
	// BEGIN __INCLUDE_EXAMPLE__
	SendSystemConfigReport(ctx context.Context, details string) error
	GetSystemConfigAll(ctx context.Context) (configs []*model.SystemConfig, err error)
	GetSystemConfigByName(ctx context.Context, name constant.SystemConfig, config interface{}) error
	UpdateSystemConfig(ctx context.Context, name constant.SystemConfig, config interface{}) error
	PublishSystemConfig(ctx context.Context, data map[constant.SystemConfig]interface{}) error
	// END __INCLUDE_EXAMPLE__
}
