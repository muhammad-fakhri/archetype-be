package httputil_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/internal/util/httputil"
	"github.com/muhammad-fakhri/go-libs/constant"
	"github.com/stretchr/testify/assert"
)

func TestGetUserID(t *testing.T) {
	type params struct {
		request func() *http.Request
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result int64, err error)
	}{
		{
			caseName: "success",
			params: params{
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					req.Header.Add(httputil.HeaderNameUserID, fmt.Sprintf("%d", testutil.UserID))
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result int64, err error) {
				assert.Nil(t, err)
				assert.Equal(t, testutil.UserID, result)
			},
		},
		{
			caseName: "missing userid",
			params: params{
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result int64, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.GetUserID(testCase.params.request())
			testCase.results(t, res, err)
		})
	}
}

func TestGetCountry(t *testing.T) {
	type params struct {
		request func() *http.Request
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result string, err error)
	}{
		{
			caseName: "success",
			params: params{
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					req.Header.Add(httputil.HeaderNameTenant, testutil.Country)
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, testutil.Country, result)
			},
		},
		{
			caseName: "missing country",
			params: params{
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.GetCountry(testCase.params.request())
			testCase.results(t, res, err)
		})
	}
}

func TestGetCmsToken(t *testing.T) {
	type params struct {
		request func() *http.Request
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result string, err error)
	}{
		{
			caseName: "success",
			params: params{
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					req.Header.Add(httputil.HeaderNameAdminToken, testutil.AdminCMSToken)
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, testutil.AdminCMSToken, result)
			},
		},
		{
			caseName: "missing token",
			params: params{
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.GetAdminToken(testCase.params.request())
			testCase.results(t, res, err)
		})
	}
}

func TestGetMktAdminToken(t *testing.T) {
	type params struct {
		request func() *http.Request
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result string, err error)
	}{
		{
			caseName: "success",
			params: params{
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					req.Header.Add(httputil.HeaderNameMktAdminToken, testutil.AdminAuthToken)
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, testutil.AdminAuthToken, result)
			},
		},
		{
			caseName: "missing token",
			params: params{
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.GetMktAdminToken(testCase.params.request())
			testCase.results(t, res, err)
		})
	}
}

func TestGetAuthCookie(t *testing.T) {
	type params struct {
		request func() []*http.Cookie
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result string, err error)
	}{
		{
			caseName: "success",
			params: params{
				request: func() []*http.Cookie {
					return []*http.Cookie{
						{
							Name:  httputil.CookieUserID,
							Value: fmt.Sprintf("%d", testutil.UserID),
						},
						{
							Name:  httputil.CookieNameUserToken,
							Value: testutil.UserAuthToken,
						},
					}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, testutil.UserAuthToken, result)
			},
		},
		{
			caseName: "missing token",
			params: params{
				request: func() []*http.Cookie {
					return []*http.Cookie{}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.GetAuthCookie(testCase.params.request(), constant.EnvTest)
			testCase.results(t, res, err)
		})
	}
}

func TestGetUserIDFromCookie(t *testing.T) {
	type params struct {
		request func() []*http.Cookie
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result int64, err error)
	}{
		{
			caseName: "success",
			params: params{
				request: func() []*http.Cookie {
					return []*http.Cookie{
						{
							Name:  httputil.CookieUserID,
							Value: fmt.Sprintf("%d", testutil.UserID),
						},
						{
							Name:  httputil.CookieNameUserToken,
							Value: testutil.UserAuthToken,
						},
					}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result int64, err error) {
				assert.Nil(t, err)
				assert.Equal(t, testutil.UserID, result)
			},
		},
		{
			caseName: "missing user id",
			params: params{
				request: func() []*http.Cookie {
					return []*http.Cookie{
						{
							Name:  httputil.CookieNameUserToken,
							Value: testutil.UserAuthToken,
						},
					}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result int64, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "failed invalid user id",
			params: params{
				request: func() []*http.Cookie {
					return []*http.Cookie{
						{
							Name:  httputil.CookieUserID,
							Value: "abc",
						},
					}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result int64, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.GetUserIDFromCookie(testCase.params.request())
			testCase.results(t, res, err)
		})
	}
}

func TestGetParamEventID(t *testing.T) {
	type params struct {
		request func() httprouter.Params
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result string, err error)
	}{
		{
			caseName: "success",
			params: params{
				request: func() httprouter.Params {
					return httprouter.Params{
						{
							Key:   httputil.ParamNameEventID,
							Value: testutil.EventID,
						},
						{
							Key:   "other param",
							Value: "other value",
						},
					}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, testutil.EventID, result)
			},
		},
		{
			caseName: "missing event id",
			params: params{
				request: func() httprouter.Params {
					return httprouter.Params{}
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.GetParamEventID(testCase.params.request())
			testCase.results(t, res, err)
		})
	}
}

func TestGetParam(t *testing.T) {
	type params struct {
		request func() *http.Request
		key     string
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params params)
		results      func(t *testing.T, result string, err error)
	}{
		{
			caseName: "success",
			params: params{
				key: httputil.HeaderNameTenant,
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					q := req.URL.Query()
					q.Add(httputil.HeaderNameTenant, testutil.Country)

					req.URL.RawQuery = q.Encode()
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, testutil.Country, result)
			},
		},
		{
			caseName: "missing value",
			params: params{
				key: httputil.HeaderNameTenant,
				request: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				},
			},
			expectations: func(params params) {},
			results: func(t *testing.T, result string, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := httputil.GetQueryParam(testCase.params.request(), testCase.params.key)
			testCase.results(t, res, err)
		})
	}
}
