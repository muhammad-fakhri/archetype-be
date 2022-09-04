package usecase_test

// BEGIN __INCLUDE_EXAMPLE__
import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/model"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
	commonconstant "github.com/muhammad-fakhri/go-libs/constant"
	"github.com/stretchr/testify/assert"
)

func TestGetSystemConfig(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	usecase := usecase.NewUsecase(m.Logger, m.Config, m.Repository)

	type params struct{}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, result *dto.PublicSystemConfig, err error)
	}{
		{
			caseName: "success",
			params:   params{},
			expectations: func(params params) {
				m.Repository.EXPECT().GetSystemConfigAll(gomock.Any()).Return([]*model.SystemConfig{
					{
						Name:   constant.CountryConfig,
						Config: "{}",
					},
					{
						Name: constant.SystemConfig("invalid-config"),
					},
				}, nil)
			},
			results: func(t *testing.T, result *dto.PublicSystemConfig, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, 1, len(result.Configs))
			},
		},
		{
			caseName: "success - invalid config",
			params:   params{},
			expectations: func(params params) {
				m.Repository.EXPECT().GetSystemConfigAll(gomock.Any()).Return([]*model.SystemConfig{
					{
						Name: constant.CountryConfig,
					},
				}, nil)
			},
			results: func(t *testing.T, result *dto.PublicSystemConfig, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, 0, len(result.Configs))
			},
		},
		{
			caseName: "failed",
			params:   params{},
			expectations: func(params params) {
				m.Repository.EXPECT().GetSystemConfigAll(gomock.Any()).Return(nil, testutil.DefaultErr)
			},
			results: func(t *testing.T, result *dto.PublicSystemConfig, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := usecase.GetSystemConfig(context.Background())
			testCase.results(t, res, err)
		})
	}
}

func TestUpdateSystemConfig(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	usecase := usecase.NewUsecase(m.Logger, m.Config, m.Repository)

	type params struct {
		in *dto.PublicSystemConfigRequest
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, err error)
	}{
		{
			caseName: "success",
			params: params{
				in: &dto.PublicSystemConfigRequest{
					Configs: map[constant.SystemConfig]interface{}{
						constant.CountryConfig: &dto.CountryConfig{
							Whitelist: []commonconstant.Country{commonconstant.ID, commonconstant.BR},
						},
					},
				},
			},
			expectations: func(params params) {
				m.Repository.EXPECT().UpdateSystemConfig(gomock.Any(), constant.CountryConfig, dto.CountryConfig{
					Whitelist: []commonconstant.Country{commonconstant.ID, commonconstant.BR},
				}).Return(nil)
				m.Repository.EXPECT().SendSystemConfigReport(gomock.Any(), string(constant.CountryConfig)).Return(nil)
			},
			results: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		{
			caseName: "failed - invalid config",
			params: params{
				in: &dto.PublicSystemConfigRequest{},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "failed - update repo",
			params: params{
				in: &dto.PublicSystemConfigRequest{
					Configs: map[constant.SystemConfig]interface{}{
						constant.CountryConfig: &dto.CountryConfig{
							Whitelist: []commonconstant.Country{commonconstant.ID, commonconstant.BR},
						},
					},
				},
			},
			expectations: func(params params) {
				m.Repository.EXPECT().UpdateSystemConfig(gomock.Any(), constant.CountryConfig, gomock.Any()).Return(testutil.DefaultErr)
			},
			results: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			testCase.results(t, usecase.UpdateSystemConfig(context.Background(), testCase.params.in))
		})
	}
}

// END __INCLUDE_EXAMPLE__
