package usecase

// BEGIN __INCLUDE_EXAMPLE__
import (
	"context"

	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/dto/converter"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
)

func (u *usecase) UpdateSystemConfig(ctx context.Context, configs *dto.PublicSystemConfigRequest) (err error) {
	configMap, err := converter.PublicToMapSystemConfig(configs)
	if err != nil {
		return
	}

	if configs.IsAsync {
		err = u.repo.PublishSystemConfig(ctx, configMap)
		if err != nil {
			u.logger.Errorf(ctx, "failed to publish update config")
			return
		}
	} else {
		err = u.UpdateSystemConfigs(ctx, configMap)
		if err != nil {
			u.logger.Errorf(ctx, "failed to update configs")
			return
		}
	}

	return
}

func (u *usecase) UpdateSystemConfigs(ctx context.Context, configMap map[constant.SystemConfig]interface{}) (err error) {
	if configMap == nil {
		return errors.Wrap(errors.ErrBadRequest, "empty request")
	}

	for name, config := range configMap {
		err = u.repo.UpdateSystemConfig(ctx, name, config)
		if err != nil {
			u.logger.Errorf(ctx, "failed to update config %s", name)
			return
		}

		// send email after update each config
		u.repo.SendSystemConfigReport(ctx, string(name))
	}

	return
}

func (u *usecase) GetSystemConfig(ctx context.Context) (configs *dto.PublicSystemConfig, err error) {
	sc, err := u.repo.GetSystemConfigAll(ctx)
	if err != nil {
		u.logger.Errorf(ctx, "failed to get configs. err:%v", err)
		return
	}

	configs = converter.ModelToPublicSystemConfig(sc)
	return
}

// END __INCLUDE_EXAMPLE__
