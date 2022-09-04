package httputil

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	json "github.com/json-iterator/go"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto/base"
)

type requestParam struct {
	param  map[string]string
	header map[string]string
	cookie map[string]string
	body   interface{}
}

type RequestOption func(*requestParam)

func WithRequestParam(param map[string]string) RequestOption {
	return func(c *requestParam) {
		c.param = param
	}
}

func WithRequestHeader(param map[string]string) RequestOption {
	return func(c *requestParam) {
		c.header = param
	}
}

func WithBody(param interface{}) RequestOption {
	return func(c *requestParam) {
		c.body = param
	}
}

func WithCookie(param map[string]string) RequestOption {
	return func(c *requestParam) {
		c.cookie = param
	}
}

func BuildBaseRequest(ctx context.Context, method, url string, opts ...RequestOption) (*http.Request, error) {
	var bodyBytes io.Reader

	opt := &requestParam{}
	for _, o := range opts {
		o(opt)
	}

	if opt.body != nil {
		marshaled, err := json.Marshal(opt.body)
		if err != nil {
			return nil, err
		}

		bodyBytes = bytes.NewReader(marshaled)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyBytes)
	if err != nil {
		return nil, err
	}

	for k, v := range opt.param {
		q := req.URL.Query()
		q.Add(k, v)
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Add("Content-Type", "application/json")
	for k, v := range opt.header {
		req.Header.Add(k, v)
	}

	for k, v := range opt.cookie {
		req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}

	return req, nil

}

func ValidateBaseResponse(respErr *base.ErrorResponse, status int) (code int, err error) {
	if respErr != nil {
		code = respErr.Code
		err = fmt.Errorf("status: %d, %s", status, respErr.Message)
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("status: %d, http error", status)
		return
	}

	return
}
