package server

import (
	"context"
	"net/http"

	"github.com/panjiesw/apimocker/db"
	"github.com/pressly/chi"
)

type Server struct {
	*chi.Mux
	DS db.Datastore
}

func New() *Server {
	// d, err := db.Open("./testdb")
	// if err != nil {
	// 	panic(err)
	// }
	r := chi.NewRouter()
	s := &Server{Mux: r}
	s.initialize()
	return s
}

func (s *Server) initialize() {
	s.Use(s.AddRootCtx)
	s.Use(s.RequestID)
	s.Use(s.LoggerMiddleware)
	s.Mount("/_admin", s.adminRouter())
}

func (s *Server) AddRootCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RootCtxKey, &RootCtx{})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
