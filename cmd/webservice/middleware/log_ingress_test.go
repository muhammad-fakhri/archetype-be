package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/middleware"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/internal/util/httputil"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto/base"
	"github.com/stretchr/testify/assert"
)

func TestEmptyMiddleware(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	mf := middleware.NewMiddlewareFactory()

	type response base.Response
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
			)

			testutil.ParseResponse(t, rr, &resp)
			testCase.results(t, &resp, rr.Code)
		})
	}
}

func TestLogIngress(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	mw := middleware.NewMiddleware(m.Usecase, m.Logger)
	mf := middleware.NewMiddlewareFactory()
	mf.Use(mw.LogIngress(
		middleware.WithoutSuccessRequest(false),
		middleware.WithoutSuccessResponseBody(false),
		middleware.WithoutIngressLog(false),
		middleware.WithoutRequestBody(false),
		middleware.WithoutRequestHeader(false),
		middleware.WithoutResponseHeader(false),
		middleware.WithoutResponseBody(false),
		middleware.WithoutRequestHeaderKeysOption([]string{httputil.HeaderNameAdminToken, httputil.HeaderNameMktAdminToken}),
	))

	type response base.Response
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
			)

			testutil.ParseResponse(t, rr, &resp)
			testCase.results(t, &resp, rr.Code)
		})
	}
}
