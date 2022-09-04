package model_test

import (
	"testing"

	"github.com/muhammad-fakhri/archetype-be/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestSystemConfigGetByName(t *testing.T) {
	type params struct{}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result []interface{})
	}{
		{
			caseName:     "success",
			params:       params{},
			expectations: func(params params) {},
			results: func(t *testing.T, result []interface{}) {
				assert.Equal(t, 1, len(result))
			},
		},
	}

	m := &model.SystemConfig{}
	for _, testCase := range testCases {
		testCase.expectations(testCase.params)
		testCase.results(t, m.GetByName())
	}
}

func TestSystemConfigGetAll(t *testing.T) {
	type params struct{}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result []interface{})
	}{
		{
			caseName:     "success",
			params:       params{},
			expectations: func(params params) {},
			results: func(t *testing.T, result []interface{}) {
				assert.Equal(t, 4, len(result))
			},
		},
	}

	m := &model.SystemConfig{}
	for _, testCase := range testCases {
		testCase.expectations(testCase.params)
		testCase.results(t, m.GetAll())
	}
}
