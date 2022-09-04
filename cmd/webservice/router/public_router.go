package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/handler"
	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/middleware"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
	"github.com/muhammad-fakhri/go-libs/log"
)

func publicRouter(router *httprouter.Router, u usecase.Usecase, l log.SLogger) {
	m := middleware.NewMiddleware(u, l)
	publicSecure := middleware.NewMiddlewareFactory()
	publicSecure.Use(m.LogIngress(
		middleware.WithoutIngressLog(true),
	))

	router.GET(PublicRoute(PingPath), handler.Ping())
}
