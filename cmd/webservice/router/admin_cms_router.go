package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/handler"
	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/middleware"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
	"github.com/muhammad-fakhri/go-libs/log"
)

func adminCmsRouter(router *httprouter.Router, u usecase.Usecase, l log.SLogger) {
	m := middleware.NewMiddleware(u, l)
	admin := middleware.NewMiddlewareFactory()
	admin.Use(m.AuthCms)
	admin.Use(m.LogIngress())
	// BEGIN __INCLUDE_EXAMPLE__
	router.GET(AdminRoute(BaseSystemConfigPath), admin.Wrap(handler.GetSystemConfig(u.GetSystemConfig)))
	router.PUT(AdminRoute(BaseSystemConfigPath), admin.Wrap(handler.UpdateSystemConfig(u.UpdateSystemConfig)))
	// END __INCLUDE_EXAMPLE__
}
