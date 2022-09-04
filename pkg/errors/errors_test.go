package errors_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/muhammad-fakhri/archetype-be/pkg/dto/base"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetErrorResponse(t *testing.T) {
	type params struct {
		err error
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result base.ErrorResponse)
	}{
		{
			caseName: "unknown error",
			params: params{
				err: errors.Wrapf(fmt.Errorf("%s", "unknown error"), "%d", 1),
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result base.ErrorResponse) {
				assert.Equal(t, http.StatusInternalServerError, result.Status)
			},
		},
		{
			caseName: "wrapped unauthorized error",
			params: params{
				err: errors.Wrap(errors.ErrAuthUnauthorized, "unknown error"),
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result base.ErrorResponse) {
				assert.Equal(t, http.StatusUnauthorized, result.Status)
			},
		},
		{
			caseName: "known error",
			params: params{
				err: errors.ErrMissingEventID,
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result base.ErrorResponse) {
				assert.Equal(t, http.StatusBadRequest, result.Status)
			},
		},
	}

	errors.EnableDebug(true)
	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			testCase.results(t, errors.GetErrorResponse(testCase.params.err))
		})
	}
}
