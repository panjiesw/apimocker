package server

import (
	"github.com/mgutz/logxi/v1"
	"github.com/panjiesw/apimocker/db"
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
	U *db.User
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "apimocker context value " + k.name
}
