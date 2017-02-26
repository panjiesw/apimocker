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
