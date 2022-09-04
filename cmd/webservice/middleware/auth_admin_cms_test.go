package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/middleware"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/internal/util/httputil"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto/base"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestAuthAdminCms(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	mw := middleware.NewMiddleware(m.Usecase, m.Logger)
	mf := middleware.NewMiddlewareFactory()
	mf.Use(mw.AuthCms)

	type response base.Response
	type params struct {
		token string
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, result *response, status int)
	}{
		{
			caseName: "failed - invalid token",
			params: params{
				token: testutil.AdminCMSToken,
			},
			expectations: func(params params) {
			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusUnauthorized, status)
			},
		},
		{
			caseName: "failed - token not exist",
			params:   params{},
			expectations: func(params params) {

			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusUnauthorized, status)
				assert.Equal(t, errors.ErrAuthMissingAuthToken.Error(), result.Errors.Message)
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
				mf.Wrap(testutil.DummyResponseWriter), http.MethodGet,
				"/",
				testutil.WithRequestHeader(map[string]string{
					httputil.HeaderNameAdminToken: testCase.params.token,
				}),
			)

			testutil.ParseResponse(t, rr, &resp)
			testCase.results(t, &resp, rr.Code)
		})
	}
}
