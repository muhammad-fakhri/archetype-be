package cryptutil_test

import (
	"encoding/base64"
	"testing"

	"github.com/muhammad-fakhri/archetype-be/internal/util/cryptutil"
	"github.com/stretchr/testify/assert"
)

func TestEncryptWithSHA256(t *testing.T) {
	type params struct {
		data string
	}

	var (
		secret = "secret"
	)
	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result string)
	}{
		{
			caseName: "check base64",
			params: params{
				data: "plain_data",
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string) {
				_, err := base64.StdEncoding.DecodeString(result)
				assert.Nil(t, err)
				assert.Greater(t, len(result), 0)
			},
		},
	}

	for _, testCase := range testCases {
		testCase.expectations(testCase.params)
		testCase.results(t, cryptutil.EncryptWithSHA256(testCase.params.data, secret))
	}
}
