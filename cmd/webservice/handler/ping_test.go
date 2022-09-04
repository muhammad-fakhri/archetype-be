package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/handler"
	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/router"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	type response struct {
		Data *dto.ResponsePing
	}
	type params struct{}

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

			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusOK, status)
				assert.Equal(t, "unit-test", result.Data.AppName)
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
				handler.Ping(), http.MethodGet,
				router.PingPath,
			)

			testutil.ParseResponse(t, rr, &resp)
			testCase.results(t, &resp, rr.Code)
		})
	}
}
