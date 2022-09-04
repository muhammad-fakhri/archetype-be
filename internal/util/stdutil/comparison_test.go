package stdutil_test

import (
	"testing"
	"time"

	"github.com/muhammad-fakhri/archetype-be/internal/util/stdutil"
	"github.com/stretchr/testify/assert"
)

func TestGetValueOrDefault(t *testing.T) {
	type params struct {
		a interface{}
		b interface{}
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result interface{})
	}{
		{
			caseName: "Compare nils",
			params: params{
				a: nil,
				b: nil,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, nil)
			},
		},
		{
			caseName: "Compare nil with empty string",
			params: params{
				a: nil,
				b: "",
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, "")
			},
		},
		{
			caseName: "Compare not nil string",
			params: params{
				a: "A",
				b: "B",
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, "A")
			},
		},
		{
			caseName: "Compare different data type (1)",
			params: params{
				a: false,
				b: "B",
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, false)
			},
		},
		{
			caseName: "Compare different data type (2)",
			params: params{
				a: 1,
				b: false,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, 1)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			testCase.results(t, stdutil.GetValueOrDefault(testCase.params.a, testCase.params.b))
		})
	}
}

func TestGetStringOrDefault(t *testing.T) {
	type params struct {
		a string
		b string
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result interface{})
	}{
		{
			caseName: "Empty string",
			params: params{
				a: "",
				b: "default",
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, "default")
			},
		},
		{
			caseName: "Non empty string",
			params: params{
				a: "exist",
				b: "default",
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, "exist")
			},
		},
		{
			caseName: "Empty default",
			params: params{
				a: "",
				b: "",
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, "")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			testCase.results(t, stdutil.GetStringOrDefault(testCase.params.a, testCase.params.b))
		})
	}
}

func TestGetIntOrDefault(t *testing.T) {
	type params struct {
		a int
		b int
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result interface{})
	}{
		{
			caseName: "Zero value",
			params: params{
				a: 0,
				b: 99,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, 99)
			},
		},
		{
			caseName: "Positive value",
			params: params{
				a: 1,
				b: 99,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, 1)
			},
		},
		{
			caseName: "Negative value",
			params: params{
				a: -1,
				b: 99,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, -1)
			},
		},
		{
			caseName: "Zero default",
			params: params{
				a: 0,
				b: 0,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, 0)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			testCase.results(t, stdutil.GetIntOrDefault(testCase.params.a, testCase.params.b))
		})
	}
}

func TestGetTimeSecondOrDefault(t *testing.T) {
	type params struct {
		a int
		b int
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result interface{})
	}{
		{
			caseName: "Zero value",
			params: params{
				a: 0,
				b: 99,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, time.Duration(99)*time.Second)
			},
		},
		{
			caseName: "Positive value",
			params: params{
				a: 1,
				b: 99,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, time.Duration(1)*time.Second)
			},
		},
		{
			caseName: "Negative value",
			params: params{
				a: -1,
				b: 99,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, time.Duration(99)*time.Second)
			},
		},
		{
			caseName: "Zero default",
			params: params{
				a: 0,
				b: 0,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result interface{}) {
				assert.Equal(t, result, time.Duration(0)*time.Second)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			testCase.results(t, stdutil.GetTimeSecondOrDefault(testCase.params.a, testCase.params.b))
		})
	}
}
