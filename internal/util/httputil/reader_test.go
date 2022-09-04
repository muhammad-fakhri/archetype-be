package httputil_test

import (
	"bytes"
	"net/http"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/internal/util/httputil"
	commonconstant "github.com/muhammad-fakhri/go-libs/constant"
	"github.com/stretchr/testify/assert"
)

type dummyConfig struct {
	Countries []commonconstant.Country `json:"countries"`
}

func TestReadRequestBody(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	type params struct {
		request func() *http.Request
		out     interface{}
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, result interface{}, err error)
	}{
		{
			caseName: "success",
			params: params{
				request: func() *http.Request {
					marshaled, _ := jsoniter.Marshal(&dummyConfig{
						Countries: []commonconstant.Country{commonconstant.MY},
					})
					bodyBytes := bytes.NewReader(marshaled)
					req, _ := http.NewRequest(http.MethodGet, "/", bodyBytes)
					return req
				},
				out: &dummyConfig{},
			},
			expectations: func(params params) {
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.Nil(t, err)
				value, ok := result.(*dummyConfig)
				assert.True(t, ok)
				assert.NotNil(t, value)
				assert.Equal(t, 1, len(value.Countries))
			},
		},
		{
			caseName: "invalid output",
			params: params{
				request: func() *http.Request {
					bodyBytes := bytes.NewReader([]byte("string"))
					req, _ := http.NewRequest(http.MethodGet, "/", bodyBytes)
					return req
				},
				out: nil,
			},
			expectations: func(params params) {
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	var err error
	for _, testCase := range testCases {
		testCase.expectations(testCase.params)

		err = httputil.ReadRequestBody(testCase.params.request(), &testCase.params.out)
		testCase.results(t, testCase.params.out, err)
	}
}
