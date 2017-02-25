package server

import (
	"github.com/pressly/chi"
)

func (s *Server) UserRoute() {
	s.Route("/users", func(r chi.Router) {

	})
}
