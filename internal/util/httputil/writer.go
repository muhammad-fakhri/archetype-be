package httputil

import (
	"net/http"

	json "github.com/json-iterator/go"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto/base"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
)

const (
	defaultOKResponse = "success"
)

func WriteResponseOK(w http.ResponseWriter, data interface{}) {
	response := &base.Response{Data: data}
	if response.Data == nil {
		response.Data = defaultOKResponse
	}

	b, status := EncodeResponse(http.StatusOK, response)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(b)
}

func WriteResponseError(w http.ResponseWriter, err error) {
	e := errors.GetErrorResponse(err)
	b, status := EncodeResponse(e.Status, &base.Response{Errors: &e})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(b)
}

func WriteResponseErrorCustom(w http.ResponseWriter, e base.ErrorResponse) {
	b, status := EncodeResponse(e.Status, &base.Response{Errors: &e})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(b)
}

func EncodeResponse(statusCode int, data interface{}) (b []byte, httpStatus int) {
	b, err := json.Marshal(data)
	if err != nil {
		e := errors.GetErrorResponse(errors.ErrEncode)
		b, _ = json.Marshal(&base.Response{Errors: &e})
		// override http status if encode failed
		httpStatus = e.Status
	} else {
		httpStatus = statusCode
	}

	return
}
