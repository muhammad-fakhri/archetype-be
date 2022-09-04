package testutil

import "net/http"

type TestHttpClient struct {
	HttpResp   *http.Response
	Resp       string
	StatusCode int
	Err        error
}

func (c *TestHttpClient) Do(request *http.Request) ([]byte, error) {
	return []byte(c.Resp), c.Err
}

func (c *TestHttpClient) DoV2(request *http.Request) ([]byte, int, error) {
	return []byte(c.Resp), c.StatusCode, c.Err
}

func (c *TestHttpClient) DoRawResponse(request *http.Request) (*http.Response, error) {
	return c.HttpResp, c.Err
}
