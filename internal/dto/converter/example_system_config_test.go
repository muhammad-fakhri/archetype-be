package converter_test

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/dto/converter"
	"github.com/muhammad-fakhri/archetype-be/internal/model"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
	commonconstant "github.com/muhammad-fakhri/go-libs/constant"
	"github.com/stretchr/testify/assert"
)

func TestModelToPublicSystemConfig(t *testing.T) {
	type params struct {
		data func() []*model.SystemConfig
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result *dto.PublicSystemConfig)
	}{
		{
			caseName: "success",
			params: params{
				data: func() []*model.SystemConfig {
					c1, _ := jsoniter.Marshal(&dto.CountryConfig{})
					return []*model.SystemConfig{
						{
							Name:   constant.CountryConfig,
							Config: string(c1),
						},
						{
							Name:   constant.SystemConfig("unknown config"),
							Config: "{}",
						},
					}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result *dto.PublicSystemConfig) {
				assert.NotNil(t, result)
				assert.Equal(t, 1, len(result.Configs))
			},
		},
		{
			caseName: "invalid model",
			params: params{
				data: func() []*model.SystemConfig {
					return []*model.SystemConfig{
						{
							Name:   constant.CountryConfig,
							Config: "invalid",
						},
					}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result *dto.PublicSystemConfig) {
				assert.NotNil(t, result)
				assert.Equal(t, 0, len(result.Configs))
			},
		},
	}

	for _, testCase := range testCases {
		testCase.expectations(testCase.params)
		testCase.results(t, converter.ModelToPublicSystemConfig(testCase.params.data()))
	}
}

func TestPublicToModelSystemConfig(t *testing.T) {
	type params struct {
		data func() *dto.PublicSystemConfigRequest
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result map[constant.SystemConfig]interface{}, err error)
	}{
		{
			caseName: "success",
			params: params{
				data: func() *dto.PublicSystemConfigRequest {
					return &dto.PublicSystemConfigRequest{
						Configs: map[constant.SystemConfig]interface{}{
							constant.CountryConfig: &dto.CountryConfig{
								Whitelist: []commonconstant.Country{commonconstant.MY},
							},
							constant.SystemConfig("other"): "",
						},
					}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result map[constant.SystemConfig]interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 1, len(result))

				var data dto.CountryConfig
				mapstructure.Decode(result[constant.CountryConfig], &data)
				assert.Equal(t, commonconstant.MY, data.Whitelist[0])
			},
		},
		{
			caseName: "invalid",
			params: params{
				data: func() *dto.PublicSystemConfigRequest {
					return &dto.PublicSystemConfigRequest{
						Configs: map[constant.SystemConfig]interface{}{
							constant.CountryConfig: "invalid",
						},
					}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result map[constant.SystemConfig]interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := converter.PublicToMapSystemConfig(testCase.params.data())
			testCase.results(t, res, err)
		})
	}
}
