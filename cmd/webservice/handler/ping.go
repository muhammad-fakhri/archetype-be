package handler

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/util/httputil"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
)

func Ping() httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		conf := config.Get()
		httputil.WriteResponseOK(rw, &dto.ResponsePing{
			Message:         "ok",
			ServerTimestamp: time.Now().Unix(),
			AppName:         conf.AppName,
			Environment:     string(conf.Environment),
		})
	}
}
