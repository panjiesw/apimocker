package server

import (
	"github.com/mgutz/logxi/v1"
)

var (
	RootCtxKey  = &contextKey{"RootContext"}
	AdminCtxKey = &contextKey{"AdminContext"}
)

type RootCtx struct {
	L  log.Logger
	ID string
}

type AdminCtx struct {
	*RootCtx
	P *Pagination
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "apimocker context value " + k.name
}
