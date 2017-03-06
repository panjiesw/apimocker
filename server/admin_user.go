package server

import (
	"net/http"

	"github.com/panjiesw/apimocker/db"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

func (s *Server) adminUserRoute() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.AdminUserList)
	r.Get("/:id", s.AdminUserByID)
	return r
}

func (s *Server) AdminUserList(w http.ResponseWriter, r *http.Request) {
	// actx := r.Context().Value(AdminCtxKey).(*AdminCtx)
}

func (s *Server) AdminUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := validateIntParam(r, "id")
	if err != nil {
		RenderAError(w, r, err)
		return
	}
	var user db.User
	if err := s.DS.UserGetByID(uint64(id), &user); err != nil {
		RenderAError(w, r, err)
	} else {
		render.JSON(w, r, &user)
	}
}
