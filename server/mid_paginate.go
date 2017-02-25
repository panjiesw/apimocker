package server

import (
	"net/http"
	"strconv"
)

func (s *Server) PaginateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var limit = 20
		var offset = 0

		query := r.URL.Query()
		limitStr := query.Get("limit")
		offsetStr := query.Get("offset")

		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = l
			}
		}

		if offsetStr != "" {
			if o, err := strconv.Atoi(offsetStr); err == nil {
				offset = o
			}
		}

		actx, _ := r.Context().Value(AdminCtxKey).(*AdminCtx)
		if actx != nil {
			actx.P = &Pagination{limit: limit, offset: offset}
		}

		next.ServeHTTP(w, r)
	})
}
