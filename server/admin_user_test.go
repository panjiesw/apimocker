package server_test

import (
	"net/http/httptest"
	"testing"

	"net/http"

	"github.com/gavv/httpexpect"
	"github.com/panjiesw/apimocker/db"
	"github.com/panjiesw/apimocker/errs"
)

type testDBAdminUser struct {
	*testDB
}

func (t *testDBAdminUser) UserGetByID(id uint64) (*db.User, *errs.AError) {
	if id == uint64(1) {
		return &db.User{Username: "foo", Email: "foo@bar.com"}, nil
	} else if id == uint64(2) {
		return nil, errs.ErrDBIDNotExists
	}
	return nil, errs.ErrDBUnkown
}

func TestServer_UserByID(t *testing.T) {
	handler := newTestServer(&testDBAdminUser{&testDB{}})
	handler.Get("/_admin/users/:id", handler.AdminUserByID)

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/_admin/users/1").Expect().
		Status(http.StatusOK).JSON().Object().
		ValueEqual("username", "foo").
		ValueEqual("email", "foo@bar.com")

	e.GET("/_admin/users/2").Expect().
		Status(http.StatusNotFound).JSON().Object().
		ValueEqual("message", errs.ErrDBIDNotExists.Msg).
		ValueEqual("type", errs.ErrDBIDNotExists.Type)

	e.GET("/_admin/users/3").Expect().
		Status(http.StatusInternalServerError).JSON().Object().
		ValueEqual("message", errs.ErrDBUnkown.Msg).
		ValueEqual("type", errs.ErrDBUnkown.Type)

	e.GET("/_admin/users/asd").Expect().
		Status(http.StatusBadRequest).JSON().Object().
		ValueEqual("message", errs.ErrRequestBadParam.Msg).
		ValueEqual("type", errs.ErrRequestBadParam.Type)
}
