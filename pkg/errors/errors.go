package errors

import (
	"errors"
	"net/http"

	"github.com/muhammad-fakhri/archetype-be/pkg/dto/base"
)

var (
	errorDebug bool
)

var (
	// server side error
	ErrUnknown           = errors.New("internal server error")
	ErrEncode            = errors.New("encode message error")
	ErrEmptyResponseData = errors.New("empty response data")
	// client side error
	ErrBadRequest                = errors.New("bad request")
	ErrMissingMandatoryParameter = errors.New("missing mandatory parameter")
	ErrMissingXUserID            = errors.New("missing x-user-id")
	ErrMissingXTenant            = errors.New("missing x-tenant")
	ErrMissingEventID            = errors.New("missing eventID")
	// access error
	ErrAuthUnauthorized     = errors.New("user is unauthorized to access this resource")
	ErrAuthMissingAuthToken = errors.New("missing auth token")
	ErrAuthInvalidUserID    = errors.New("invalid user id")
	ErrTooManyRequest       = errors.New("too many request")
	/* internal server error detail
	by default will be remapped to ErrUnknown, unwrap with errors.Is if client need to know detail root cause*/
	ErrDatabase  = errors.New("database error")
	ErrRedis     = errors.New("redis error")
	ErrHost      = errors.New("host error")
	ErrHostEmail = errors.New("email error")

	// service specific error
	// ...
)

type ErrCode int

/*
	server side error: 9xxx
	client side error: 4xxx
	security error: 3xxx
	application specific business logic error: 1xxx - 2xxx
	undefined: 0
*/
const (
	ErrUndefined                     ErrCode = 0
	ErrCodeUnknown                   ErrCode = 9999 // unknown server side error
	ErrCodeEncode                    ErrCode = 9000 // response encoding error
	ErrCodeEmptyResponseData         ErrCode = 9001
	ErrCodeBadRequest                ErrCode = 4999
	ErrCodeMissingMandatoryParameter ErrCode = 4000
	ErrCodeMissingXUserID            ErrCode = 4001
	ErrCodeMissingXTenant            ErrCode = 4002
	ErrCodeMissingEventID            ErrCode = 4003
	ErrCodeAuthUnauthorized          ErrCode = 3999
	ErrCodeAuthInvalidUserID         ErrCode = 3000
	ErrCodeMissingAuthToken          ErrCode = 3001
	ErrCodeTooManyRequest            ErrCode = 3002
	ErrCodeAggregatorNotSupported    ErrCode = 1000
	ErrCodeAggregatorLinkEmpty       ErrCode = 1001
	ErrCodeResourceNotFound          ErrCode = 1002
)

const (
	ErrStatServer   = http.StatusInternalServerError
	ErrStatClient   = http.StatusBadRequest
	ErrStatAuth     = http.StatusUnauthorized
	ErrStatNotFound = http.StatusNotFound
)

var errorMap = map[error]base.ErrorResponse{
	ErrUnknown:                   ErrorResponse(ErrStatServer, ErrCodeUnknown, ErrUnknown),
	ErrEncode:                    ErrorResponse(ErrStatServer, ErrCodeUnknown, ErrEncode),
	ErrEmptyResponseData:         ErrorResponse(ErrStatServer, ErrCodeEmptyResponseData, ErrEmptyResponseData),
	ErrBadRequest:                ErrorResponse(ErrStatClient, ErrCodeBadRequest, ErrBadRequest),
	ErrMissingMandatoryParameter: ErrorResponse(ErrStatClient, ErrCodeMissingMandatoryParameter, ErrMissingMandatoryParameter),
	ErrMissingXTenant:            ErrorResponse(ErrStatClient, ErrCodeMissingXTenant, ErrMissingXTenant),
	ErrMissingXUserID:            ErrorResponse(ErrStatClient, ErrCodeMissingXUserID, ErrMissingXUserID),
	ErrMissingEventID:            ErrorResponse(ErrStatClient, ErrCodeMissingEventID, ErrMissingEventID),
	ErrAuthUnauthorized:          ErrorResponse(ErrStatAuth, ErrCodeAuthUnauthorized, ErrAuthUnauthorized),
	ErrAuthInvalidUserID:         ErrorResponse(ErrStatAuth, ErrCodeAuthInvalidUserID, ErrAuthInvalidUserID),
	ErrAuthMissingAuthToken:      ErrorResponse(http.StatusUnauthorized, ErrCodeMissingAuthToken, ErrAuthMissingAuthToken),
	ErrTooManyRequest:            ErrorResponse(http.StatusTooManyRequests, ErrCodeTooManyRequest, ErrTooManyRequest),
}

func EnableDebug(flag bool) {
	errorDebug = flag
}

func ErrorResponse(status int, code ErrCode, err error) base.ErrorResponse {
	return base.ErrorResponse{
		Status:  status,
		Code:    int(code),
		Message: err.Error(),
	}
}

func GetErrorResponse(err error) (er base.ErrorResponse) {
	er, ok := errorMap[err]
	if !ok {
		if errors.Is(err, ErrAuthUnauthorized) {
			er = ErrorResponse(ErrStatAuth, ErrCodeAuthUnauthorized, err)
		} else if errors.Is(err, ErrBadRequest) {
			er = ErrorResponse(ErrStatClient, ErrCodeBadRequest, err)
		} else {
			er = ErrorResponse(ErrStatServer, ErrCodeUnknown, err)
		}

		if errorDebug {
			er.Detail = err.Error()
		}
	}

	return
}
