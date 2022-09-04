package cacheutil_test

// BEGIN __INCLUDE_REDIS__
import (
	"math"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/internal/util/cacheutil"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var jsonr = jsoniter.ConfigCompatibleWithStandardLibrary

func TestCacheGetJSON(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	type params struct {
		key   string
		out   interface{}
		value func() string
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result interface{}, err error)
	}{
		{
			caseName: "success",
			params: params{
				key: "cache-key",
				value: func() string {
					value, _ := jsoniter.Marshal(map[string]float64{
						"key_1": 1,
						"key_2": 2,
					})
					return string(value)
				},
			},
			expectations: func(params params) {
				m.Cache.EXPECT().Get(params.key).Return(params.value(), nil)
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.Nil(t, err)
				value, ok := result.(map[string]interface{})
				assert.True(t, ok)
				assert.NotNil(t, value)
				assert.Equal(t, 2, len(value))
			},
		},
		{
			caseName: "empty",
			params: params{
				key: "cache-key",
			},
			expectations: func(params params) {
				m.Cache.EXPECT().Get(params.key).Return("", nil)
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "error",
			params: params{
				key: "cache-key",
			},
			expectations: func(params params) {
				m.Cache.EXPECT().Get(params.key).Return("", errors.ErrUnknown)
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			err := cacheutil.GetJSON(testCase.params.key, &testCase.params.out, m.Cache.Get)
			testCase.results(t, testCase.params.out, err)
		})
	}
}

func TestCacheSetJSON(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	type params struct {
		key   string
		value interface{}
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, err error)
	}{
		{
			caseName: "success",
			params: params{
				key: "cache-key",
				value: map[string]float64{
					"key_1": 1,
					"key_2": 2,
				},
			},
			expectations: func(params params) {
				value, _ := jsonr.Marshal(params.value)
				m.Cache.EXPECT().Set(params.key, string(value), 1*time.Minute).Return(nil)
			},
			results: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		{
			caseName: "error",
			params: params{
				key:   "cache-key",
				value: math.Inf(1),
			},
			expectations: func(params params) {

			},
			results: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			testCase.results(t, cacheutil.SetJSON(testCase.params.key, testCase.params.value, 1*time.Minute, m.Cache.Set))
		})
	}
}

// END __INCLUDE_REDIS__
