package server

import (
	"net/http"
	"strconv"

	"github.com/panjiesw/apimocker/db"
)

func (s *Server) PaginateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var limit = uint(20)
		var offset = uint(0)

		query := r.URL.Query()
		limitStr := query.Get("limit")
		offsetStr := query.Get("offset")

		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = uint(l)
			}
		}

		if offsetStr != "" {
			if o, err := strconv.Atoi(offsetStr); err == nil {
				offset = uint(o)
			}
		}

		actx, _ := r.Context().Value(AdminCtxKey).(*AdminCtx)
		if actx != nil {
			actx.M = &db.Meta{Limit: limit, Offset: offset}
		}

		next.ServeHTTP(w, r)
	})
}
