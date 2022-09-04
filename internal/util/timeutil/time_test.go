package timeutil_test

import (
	"testing"
	"time"

	"github.com/muhammad-fakhri/archetype-be/internal/util/timeutil"
	"github.com/stretchr/testify/assert"
)

func TestGetTimeMillis(t *testing.T) {
	type params struct{}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result int64)
	}{
		{
			caseName:     "success",
			params:       params{},
			expectations: func(params params) {},
			results: func(t *testing.T, result int64) {
				assert.LessOrEqual(t, result, time.Now().UnixNano()/int64(time.Millisecond))
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			testCase.results(t, timeutil.NowMillis())
		})
	}
}

func TestGetTimeSecond(t *testing.T) {
	type params struct{}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result int64)
	}{
		{
			caseName:     "success",
			params:       params{},
			expectations: func(params params) {},
			results: func(t *testing.T, result int64) {
				assert.LessOrEqual(t, result, time.Now().UnixNano()/int64(time.Second))
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			testCase.results(t, timeutil.NowSeconds())
		})
	}
}
