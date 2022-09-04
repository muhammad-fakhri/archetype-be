package handler

// BEGIN __INCLUDE_EXAMPLE__
import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/internal/util/httputil"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
)

type GetSystemConfigHandler func(ctx context.Context) (configs *dto.PublicSystemConfig, err error)

func GetSystemConfig(f GetSystemConfigHandler) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		res, err := f(r.Context())
		if err != nil {
			httputil.WriteResponseError(rw, err)
			return
		}

		httputil.WriteResponseOK(rw, res)
	}
}

type UpdateSystemConfigHandler func(ctx context.Context, configs *dto.PublicSystemConfigRequest) (err error)

func UpdateSystemConfig(f UpdateSystemConfigHandler) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var req dto.PublicSystemConfigRequest
		err := httputil.ReadRequestBody(r, &req)
		if err != nil {
			httputil.WriteResponseError(rw, errors.ErrBadRequest)
			return
		}

		err = f(r.Context(), &req)
		if err != nil {
			httputil.WriteResponseError(rw, err)
			return
		}

		httputil.WriteResponseOK(rw, nil)
	}
}

// END __INCLUDE_EXAMPLE__
