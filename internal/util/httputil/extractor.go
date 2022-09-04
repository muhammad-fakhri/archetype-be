package httputil

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	commonconstant "github.com/muhammad-fakhri/go-libs/constant"
)

const (
	// header
	HeaderNameTenant        = "x-tenant"
	HeaderNameUserID        = "x-user-id"
	HeaderNameMktAdminToken = "x-mkt-token"
	HeaderNameAdminToken    = "Authorization"

	// param
	ParamNameEventID = "eventID"

	// cookie
	CookieUserID        = "SPC_U"
	CookieNameUserToken = "SPC_ST"
	CookieNameMktAuth   = "MKT_SSO_MSG"
)

func GetParamEventID(ps httprouter.Params) (eventID string, err error) {
	eventID = ps.ByName(ParamNameEventID)
	if eventID == "" {
		err = errors.ErrMissingEventID
	}

	return
}

func GetUserID(r *http.Request) (int64, error) {
	userID, err := strconv.ParseInt(r.Header.Get(HeaderNameUserID), 10, 64)
	if err != nil || userID <= 0 {
		return userID, errors.ErrMissingXUserID
	}

	return userID, nil
}

func GetUserIDFromCookie(cookies []*http.Cookie) (int64, error) {
	for _, c := range cookies {
		if c.Name != CookieUserID {
			continue
		}

		userID, err := strconv.Atoi(c.Value)
		if err != nil {
			return 0, err
		}

		return int64(userID), nil
	}

	return 0, errors.ErrMissingXUserID
}

func GetCountry(r *http.Request) (string, error) {
	country := strings.ToUpper(r.Header.Get(HeaderNameTenant))
	if country == "" {
		return "", errors.ErrMissingXTenant
	}

	return country, nil
}

func GetMktAdminToken(r *http.Request) (string, error) {
	token := r.Header.Get(HeaderNameMktAdminToken)
	if token == "" {
		return "", errors.ErrAuthMissingAuthToken
	}

	return token, nil
}

func GetAdminToken(r *http.Request) (string, error) {
	token := r.Header.Get(HeaderNameAdminToken)
	if token == "" {
		return "", errors.ErrAuthMissingAuthToken
	}

	return token, nil
}

func GetAuthCookie(cookies []*http.Cookie, env commonconstant.Environment) (token string, err error) {
	for _, c := range cookies {
		if c.Name == CookieNameUserToken && c.Value != "" && c.Value != "-" &&
			(strings.Contains(c.Domain, constant.DomainEnv[env]) || c.Domain == "") {
			token = c.Value
			break
		}
	}

	if len(token) == 0 {
		err = errors.ErrAuthMissingAuthToken
	}

	return
}

func GetQueryParam(r *http.Request, key string) (value string, err error) {
	queryValues := r.URL.Query()
	value = queryValues.Get(key)
	if len(value) == 0 {
		err = errors.ErrMissingMandatoryParameter
		return
	}

	return value, nil
}
