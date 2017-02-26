package server

import (
	"net/http"
	"strconv"

	"github.com/panjiesw/apimocker/errs"
	"github.com/pressly/chi"
)

func validateIntParam(r *http.Request, name string) (result int, err *errs.AError) {
	result, e := strconv.Atoi(chi.URLParam(r, name))
	if e != nil {
		err = errs.ErrRequestBadParam
	}
	return
}
