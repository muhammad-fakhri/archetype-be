package middleware

import (
	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
	"github.com/muhammad-fakhri/go-libs/constant"
	"github.com/muhammad-fakhri/go-libs/log"
)

type middleware struct {
	usecase usecase.Usecase
	logger  log.SLogger
	config  middlewareConfig
}

type middlewareConfig struct {
	jwtSecret string
	env       constant.Environment
}

type Middleware interface {
	// middleware functions
	AuthCms(next httprouter.Handle) httprouter.Handle
	LogIngress(opts ...LogOption) func(next httprouter.Handle) httprouter.Handle
}

func NewMiddleware(u usecase.Usecase, l log.SLogger) Middleware {
	conf := config.Get()
	return &middleware{
		usecase: u,
		logger:  l,
		config: middlewareConfig{
			jwtSecret: conf.AuthConfig.JWTSecret,
			env:       conf.Environment,
		},
	}
}

type middlewareFactory struct {
	middlewares []MiddlewareFunc
}

type MiddlewareFactory interface {
	// Use will append middleware to request wrapper list.
	// Note: Middleware insert position matters! any Middleware that pushed last will be executed first (i.e. Middleware
	// will be executed in reverse order)
	Use(...MiddlewareFunc)
	Wrap(httprouter.Handle) httprouter.Handle
}

func NewMiddlewareFactory() MiddlewareFactory {
	return &middlewareFactory{}
}

// MiddlewareFunc represents middleware function type
type MiddlewareFunc func(httprouter.Handle) httprouter.Handle

func (m *middlewareFactory) Use(mw ...MiddlewareFunc) {
	m.middlewares = append(m.middlewares, mw...)
}

func (m *middlewareFactory) Wrap(fn httprouter.Handle) httprouter.Handle {
	l := len(m.middlewares)
	if l == 0 {
		return fn
	}

	result := fn
	for _, mw := range m.middlewares {
		result = mw(result)
	}

	return result
}
