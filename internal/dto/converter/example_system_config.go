package converter

// BEGIN __INCLUDE_EXAMPLE__
import (
	"log"

	json "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"

	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/model"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
)

func ModelToPublicSystemConfig(in []*model.SystemConfig) (out *dto.PublicSystemConfig) {
	out = &dto.PublicSystemConfig{
		Configs: make(map[constant.SystemConfig]*dto.SystemConfig),
	}

	for _, val := range in {
		c := &dto.SystemConfig{
			UpdatedAt: val.UpdatedAt,
			CreatedAt: val.CreatedAt,
		}
		switch val.Name {
		case constant.CountryConfig:
			c.Config = &dto.CountryConfig{}
			err := json.Unmarshal([]byte(val.Config), &c.Config)
			if err != nil {
				log.Printf("invalid config %s, value %s. err:%v", val.Name, val.Config, err)
				continue
			}
			out.Configs[val.Name] = c
		default:
			log.Printf("unknown config type %s", val.Name)
			continue
		}
	}

	return
}

func PublicToMapSystemConfig(in *dto.PublicSystemConfigRequest) (out map[constant.SystemConfig]interface{}, err error) {
	out = make(map[constant.SystemConfig]interface{}, 0)

	for key, val := range in.Configs {
		switch key {
		case constant.CountryConfig:
			var config dto.CountryConfig
			mapstructure.Decode(val, &config)

			if config.Whitelist == nil {
				return nil, errors.Wrap(errors.ErrBadRequest, "empty country config")
			}

			out[key] = config
		default:
			continue
		}
	}

	if len(out) == 0 {
		return nil, errors.Wrap(errors.ErrBadRequest, "empty config")
	}

	return
}

// END __INCLUDE_EXAMPLE__
