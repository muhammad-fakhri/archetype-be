package handler_test

// BEGIN __INCLUDE_EXAMPLE__
import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/handler"
	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/router"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto/base"
	"github.com/stretchr/testify/assert"
)

func TestGetSystemConfig(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	type response base.Response
	type params struct {
		Path string
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, result *response, status int)
	}{
		{
			caseName: "success",
			params:   params{},
			expectations: func(params params) {
				m.Usecase.EXPECT().GetSystemConfig(gomock.Any()).Return(&dto.PublicSystemConfig{}, nil)
			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusOK, status)
			},
		},
		{
			caseName: "failed - error usecase",
			params:   params{},
			expectations: func(params params) {
				m.Usecase.EXPECT().GetSystemConfig(gomock.Any()).Return(nil, testutil.DefaultErr)
			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusInternalServerError, status)
			},
		},
	}

	var (
		resp response
		rr   *httptest.ResponseRecorder
	)

	for _, testCase := range testCases {
		testCase.expectations(testCase.params)

		rr = testutil.NewRequestRecorder(t,
			handler.GetSystemConfig(m.Usecase.GetSystemConfig), http.MethodGet,
			testutil.GetRouterPath(testCase.params.Path, router.BaseSystemConfigPath),
		)

		testutil.ParseResponse(t, rr, &resp)
		testCase.results(t, &resp, rr.Code)
	}
}

func TestUpdateSystemConfig(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	type response base.Response
	type params struct {
		Path string
		Body interface{}
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, result *response, status int)
	}{
		{
			caseName: "success",
			params: params{
				Body: &dto.PublicSystemConfigRequest{},
			},
			expectations: func(params params) {
				m.Usecase.EXPECT().UpdateSystemConfig(gomock.Any(), params.Body).Return(nil)
			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusOK, status)
			},
		},
		{
			caseName: "failed - error usecase",
			params: params{
				Body: &dto.PublicSystemConfigRequest{},
			},
			expectations: func(params params) {
				m.Usecase.EXPECT().UpdateSystemConfig(gomock.Any(), params.Body).Return(testutil.DefaultErr)
			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusInternalServerError, status)
			},
		},
		{
			caseName: "failed - invalid request body",
			params: params{
				Body: "invalid data",
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusBadRequest, status)
			},
		},
	}

	var (
		resp response
		rr   *httptest.ResponseRecorder
	)

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)

			rr = testutil.NewRequestRecorder(t,
				handler.UpdateSystemConfig(m.Usecase.UpdateSystemConfig), http.MethodPut,
				testutil.GetRouterPath(testCase.params.Path, router.BaseSystemConfigPath),
				testutil.WithBody(testCase.params.Body),
			)

			testutil.ParseResponse(t, rr, &resp)
			testCase.results(t, &resp, rr.Code)
		})
	}
}

// END __INCLUDE_EXAMPLE__
