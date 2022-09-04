package httputil_test

import (
	"context"
	"math"
	"net/http"
	"strings"
	"testing"

	"github.com/muhammad-fakhri/archetype-be/internal/util/httputil"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto/base"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestBuildBaseRequest(t *testing.T) {
	type params struct {
		method string
		url    string
		body   interface{}
		cookie map[string]string
		header map[string]string
		param  map[string]string
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result *http.Request, err error)
	}{
		{
			caseName: "success",
			params: params{
				method: http.MethodGet,
				url:    "http://test-host/test-api",
				body: map[string]interface{}{
					"id": 1,
				},
				cookie: map[string]string{
					"c1": "cookie_1",
				},
				header: map[string]string{
					"h1": "header_1",
				},
				param: map[string]string{
					"p1": "param_1",
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result *http.Request, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, http.MethodGet, result.Method)
				assert.Equal(t, "test-host", result.URL.Host)
				assert.Equal(t, "/test-api", result.URL.Path)
			},
		},
		{
			caseName: "invalid body",
			params: params{
				method: http.MethodGet,
				url:    "http://test-host/test-api",
				body:   math.Inf(1),
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result *http.Request, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, result)
			},
		},
		{
			caseName: "invalid url",
			params: params{
				url: ":",
				body: map[string]interface{}{
					"id": 1,
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result *http.Request, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.BuildBaseRequest(context.Background(), testCase.params.method, testCase.params.url,
				httputil.WithBody(testCase.params.body),
				httputil.WithCookie(testCase.params.cookie),
				httputil.WithRequestHeader(testCase.params.header),
				httputil.WithRequestParam(testCase.params.param),
			)
			testCase.results(t, res, err)
		})
	}
}

func TestValidateBaseResponse(t *testing.T) {
	type params struct {
		err    *base.ErrorResponse
		status int
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, code int, err error)
	}{
		{
			caseName: "success",
			params: params{
				err:    nil,
				status: http.StatusOK,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, code int, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int(errors.ErrUndefined), code)
			},
		},
		{
			caseName: "negative status",
			params: params{
				err:    nil,
				status: http.StatusBadRequest,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, code int, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, int(errors.ErrUndefined), code)
			},
		},
		{
			caseName: "business logic error",
			params: params{
				err: &base.ErrorResponse{
					Code:    int(errors.ErrCodeAuthUnauthorized),
					Message: errors.ErrAuthUnauthorized.Error(),
				},
				status: http.StatusUnauthorized,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, code int, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, int(errors.ErrCodeAuthUnauthorized), code)
				assert.True(t, strings.Contains(err.Error(), errors.ErrAuthUnauthorized.Error()))
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.ValidateBaseResponse(testCase.params.err, testCase.params.status)
			testCase.results(t, res, err)
		})
	}
}
