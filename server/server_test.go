package server_test

import (
	"context"
	"net/http"

	"github.com/panjiesw/apimocker/db"
	"github.com/panjiesw/apimocker/errs"
	"github.com/panjiesw/apimocker/server"
	"github.com/pressly/chi"
)

type testDB struct {
}

func (t *testDB) UserSave(u *db.User) *errs.AError {
	return nil
}
func (t *testDB) UserUsernameExist(username string) (bool, *errs.AError) {
	return false, nil
}
func (t *testDB) UserEmailExist(email string) (bool, *errs.AError) {
	return false, nil
}
func (t *testDB) UserGetByUsername(username string) (*db.User, *errs.AError) {
	return nil, nil
}
func (t *testDB) UserGetByEmail(email string) (*db.User, *errs.AError) {
	return nil, nil
}
func (t *testDB) UserGetByID(id uint64) (*db.User, *errs.AError) {
	return nil, nil
}

func addTestRootCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), server.RootCtxKey, &server.RootCtx{})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func addTestAdminCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := r.Context().Value(server.RootCtxKey).(*server.RootCtx)
		ctx := context.WithValue(r.Context(), server.AdminCtxKey, &server.AdminCtx{RootCtx: rctx})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func newTestServer(ds db.Datastore) *server.Server {
	if ds == nil {
		ds = &testDB{}
	}
	r := chi.NewRouter()
	s := &server.Server{Mux: r, DS: ds}
	s.Use(addTestRootCtx)
	s.Use(addTestAdminCtx)
	return s
}
