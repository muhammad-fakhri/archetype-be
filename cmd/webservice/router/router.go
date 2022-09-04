package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
	"github.com/muhammad-fakhri/go-libs/log"
	"github.com/rs/cors"
)

func Init(u usecase.Usecase, logger log.SLogger) http.Handler {
	router := httprouter.New()
	publicRouter(router, u, logger)
	adminCmsRouter(router, u, logger)
	adminMktRouter(router, u, logger)

	corsLib := cors.New(cors.Options{
		AllowedOrigins: GetAllowedOrigins(string(config.Get().Environment)),
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "X-Tenant", "X-User-ID"},
	})

	return corsLib.Handler(router)
}

func GetAllowedOrigins(env string) []string {
	origins := []string{"http://localhost*"} // debug purposes

	switch env {
	case "production":
		env = ""
	default:
		origins = append(origins, "http://192*", "http://10*") // FE debug purposes
	}

	return origins
}
