package server

import (
	"net/http"

	"context"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

type Pagination struct {
	limit  int
	offset int
}

func (s *Server) AddAdminCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := r.Context().Value(RootCtxKey).(*RootCtx)
		ctx := context.WithValue(r.Context(), AdminCtxKey, &AdminCtx{RootCtx: rctx})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) adminRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(s.AddAdminCtx)
	// r.Use(s.AuthMiddleware)
	r.Use(s.PaginateMiddleware)
	r.Get("/", s.Version)
	return r
}

func (s *Server) Version(w http.ResponseWriter, r *http.Request) {
	actx := r.Context().Value(AdminCtxKey).(*AdminCtx)
	render.JSON(w, r, map[string]interface{}{"version": "1.0.0", "limit": actx.P.limit})
}
