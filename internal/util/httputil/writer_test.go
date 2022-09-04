package httputil_test

import (
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/internal/util/httputil"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto/base"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	type response base.Response
	type params struct {
		handler func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params)
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, result *response, status int)
	}{
		{
			caseName: "ok",
			params: params{
				handler: func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
					httputil.WriteResponseOK(w, "success")
				},
			},
			expectations: func(params params) {
			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusOK, status)
			},
		},
		{
			caseName: "ok",
			params: params{
				handler: func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
					httputil.WriteResponseOK(w, nil)
				},
			},
			expectations: func(params params) {
			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusOK, status)
			},
		},
		{
			caseName: "error",
			params: params{
				handler: func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
					httputil.WriteResponseError(w, errors.ErrAuthInvalidUserID)
				},
			},
			expectations: func(params params) {
			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusUnauthorized, status)
			},
		},
		{
			caseName: "error custom",
			params: params{
				handler: func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
					httputil.WriteResponseErrorCustom(w, errors.ErrorResponse(http.StatusInternalServerError, 1000, fmt.Errorf("custom")))
				},
			},
			expectations: func(params params) {
			},
			results: func(t *testing.T, result *response, status int) {
				assert.Equal(t, http.StatusInternalServerError, status)
			},
		},
		{
			caseName: "invalid encode",
			params: params{
				handler: func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
					httputil.WriteResponseOK(w, math.Inf(1))
				},
			},
			expectations: func(params params) {
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
			testCase.params.handler, http.MethodGet,
			"/",
		)

		testutil.ParseResponse(t, rr, &resp)
		testCase.results(t, &resp, rr.Code)
	}
}
