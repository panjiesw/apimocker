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
		var err *errs.AError

		actx := r.Context().Value(AdminCtxKey).(*AdminCtx)

		j, e := jws.ParseJWTFromRequest(r)
		if e != nil {
			actx.L.Error("failed to fetch jws", "error", e)
			err = errs.ErrAuthNoToken
		}

		if j != nil {
			v := jws.NewValidator(ex, time.Duration(24)*time.Hour, 0, nil)
			if e = j.Validate("somekey", crypto.SigningMethodHS512, v); e != nil {
				actx.L.Error("failed to validate token", "error", e)
				err = errs.ErrAuthInvalidToken
			} else {
				if subject, ok := j.Claims().Subject(); ok && subject != "" {
					if u, e := s.DS.UserGetByUsername(subject); e != nil {
						actx.L.Error("failed to find user", "error", e)
						switch ee := e.(type) {
						case *errs.AError:
							err = ee
						default:
							err = errs.ErrDBUnkown
						}
					} else {
						actx.U = u
					}
				} else {
					actx.L.Error("failed to subject", "subject", subject, "ok", ok)
					err = errs.ErrAuthInvalidToken
				}
			}
		}

		if err != nil {
			render.Status(r, err.Status)
			render.Respond(w, r, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
