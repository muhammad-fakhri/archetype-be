package middleware

import (
	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/go-libs/httpmiddleware"
)

func (m *middleware) LogIngress(opts ...LogOption) func(next httprouter.Handle) httprouter.Handle {
	// default
	options := httpmiddleware.Config{
		ExcludeOpt: &httpmiddleware.ExcludeOption{
			RequestHeader:     true,
			ResponseHeader:    true,
			RequestHeaderKeys: []string{},
		},
		DisableIngressLog: false,
	}

	for _, o := range opts {
		o(&options)
	}

	logIngressMiddleware := httpmiddleware.NewIngressLogMiddleware(
		m.logger,
		&options,
	)

	return logIngressMiddleware.EnforceWithParams
}

type LogOption func(*httpmiddleware.Config)

// returns the option to exclude response header from log
func WithoutRequestHeader(param bool) LogOption {
	return func(opts *httpmiddleware.Config) {
		opts.ExcludeOpt.RequestHeader = param
	}
}

//  returns the option to exclude request body from log
func WithoutRequestBody(param bool) LogOption {
	return func(opts *httpmiddleware.Config) {
		opts.ExcludeOpt.RequestBody = param
	}
}

// returns the option to exclude response header from log
func WithoutResponseHeader(param bool) LogOption {
	return func(opts *httpmiddleware.Config) {
		opts.ExcludeOpt.ResponseHeader = param
	}
}

// returns the option to exclude response body from log
func WithoutResponseBody(param bool) LogOption {
	return func(opts *httpmiddleware.Config) {
		opts.ExcludeOpt.ResponseBody = param
	}
}

// returns the option to exclude success response body from log
func WithoutSuccessResponseBody(param bool) LogOption {
	return func(opts *httpmiddleware.Config) {
		opts.ExcludeOpt.SuccessResponseBody = param
	}
}

// returns the option to exclude success request from log
func WithoutSuccessRequest(param bool) LogOption {
	return func(opts *httpmiddleware.Config) {
		opts.ExcludeOpt.SuccessRequest = param
	}
}

// returns the option to exclude specific header from log
func WithoutRequestHeaderKeysOption(param []string) LogOption {
	return func(opts *httpmiddleware.Config) {
		opts.ExcludeOpt.RequestHeaderKeys = param
	}
}

// returns the option to exclude ingress log entirely while appending mandatory field only (e.g: user_id, context_id, event_id, country)
func WithoutIngressLog(param bool) LogOption {
	return func(opts *httpmiddleware.Config) {
		opts.DisableIngressLog = param
	}
}
