package middleware

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/util/httputil"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/muhammad-fakhri/go-libs/authz/jwtauthz"
)

func (m *middleware) AuthCms(next httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token, err := httputil.GetAdminToken(r)
		if err != nil {
			httputil.WriteResponseError(rw, err)
			return
		}

		auth := jwtauthz.NewJwtAuthorization(m.config.jwtSecret, token)

		err = auth.ExtractClaimsFromToken()
		if err != nil {
			httputil.WriteResponseError(rw, errors.ErrAuthUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), constant.AuthorizedAdminIDKey, auth.JwtClaims.Email)
		next(rw, r.WithContext(ctx), ps)

	}
}
