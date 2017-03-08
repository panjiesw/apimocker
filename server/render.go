package server

import (
	"net/http"

	"github.com/panjiesw/apimocker/errs"
	"github.com/pressly/chi/render"
)

func RenderAError(w http.ResponseWriter, r *http.Request, err *errs.AError) {
	render.Status(r, err.Status)
	render.JSON(w, r, err)
}

func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	var ee *errs.AError
	switch e := err.(type) {
	case *errs.AError:
		ee = e
	default:
		ee = errs.New("server", e.Error(), http.StatusInternalServerError)
	}
	RenderAError(w, r, ee)
}
