package server

import (
	"net/http"

	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/panjiesw/apimocker/errs"
	"github.com/pressly/chi/render"
)

func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	ex := jws.Claims{}
	ex.SetAudience("http://localhost:3000")
	ex.SetIssuer("apimocker")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actx := r.Context().Value(AdminCtxKey).(*AdminCtx)

		if err := s.parseAndValidateToken(actx, r, ex); err != nil {
			render.Status(r, err.Status)
			render.Respond(w, r, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) parseAndValidateToken(actx *AdminCtx, r *http.Request, ex jws.Claims) *errs.AError {
	j, e := jws.ParseJWTFromRequest(r)
	if e != nil {
		actx.L.Error("failed to fetch jws", "error", e)
		return errs.ErrAuthNoToken
	}

	if j != nil {
		v := jws.NewValidator(ex, time.Duration(24)*7*time.Hour, 0, nil)
		if e = j.Validate("somekey", crypto.SigningMethodHS512, v); e != nil {
			actx.L.Error("failed to validate token", "error", e)
			return errs.ErrAuthInvalidToken
		}

		if subject, ok := j.Claims().Subject(); ok && subject != "" {
			u, err := s.DS.UserGetByUsername(subject)
			if err != nil {
				return err
			}
			actx.U = u
		} else {
			actx.L.Error("failed to subject", "subject", subject, "ok", ok)
			return errs.ErrAuthInvalidToken
		}
	}
	return nil
}
