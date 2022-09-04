package webservice

import (
	"context"
	"log"
	"net/http"

	"github.com/muhammad-fakhri/archetype-be/cmd/webservice/router"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
	logger "github.com/muhammad-fakhri/go-libs/log"
)

var (
	srv *http.Server
)

// Start initiate the webservice binary
func Start() (stopFunc func()) {
	conf := config.Get()

	logger := logger.NewSLogger(conf.AppName)
	usecase := usecase.InitDependencies(logger)
	router := router.Init(usecase, logger)

	srv = &http.Server{
		Addr:    ":" + conf.HttpPort,
		Handler: router,
	}
	go func() {
		log.Printf("running %s in web service mode...\n", conf.AppName)
		log.Println("starting web, listening on", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln("failed starting web on", srv.Addr, err)
		}
	}()
	return func() {
		GracefulStop()
	}
}

// GracefulStop stop the web service
func GracefulStop() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.WebGracefulStopTimeout)
	defer cancel()

	log.Println("shuting down web on", srv.Addr)
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatalln("failed shutdown server", err)
	}
	log.Println("web gracefully stopped")
	return
}
