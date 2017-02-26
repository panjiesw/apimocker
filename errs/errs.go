package errs

import (
	"fmt"
	"net/http"
)

type AError struct {
	Type   string `json:"type"`
	Msg    string `json:"message"`
	Status int    `json:"-"`
}

func (e *AError) Error() string {
	return fmt.Sprintf("%s - %s", e.Type, e.Msg)
}

func New(typ, msg string, status int) *AError {
	return &AError{Type: typ, Msg: msg, Status: status}
}

var (
	// ErrDBUsernameExists is returned when trying to signup using existing username.
	ErrDBUsernameExists    = &AError{Type: "db", Msg: "username already exists", Status: http.StatusNotFound}
	ErrDBUsernameNotExists = &AError{Type: "db", Msg: "username doesn't exist", Status: http.StatusNotFound}
	// ErrDBEmailExists is returned when trying to signup using existing email.
	ErrDBEmailExists    = &AError{Type: "db", Msg: "email already exists", Status: http.StatusBadRequest}
	ErrDBEmailNotExists = &AError{Type: "db", Msg: "email doesn't exist", Status: http.StatusNotFound}
	ErrDBIDNotExists    = &AError{Type: "db", Msg: "id doesn't exist", Status: http.StatusNotFound}

	ErrDBUnkown = &AError{Type: "db", Msg: "internal error", Status: http.StatusInternalServerError}

	ErrAuthNoToken      = &AError{Type: "auth", Msg: "no token", Status: http.StatusUnauthorized}
	ErrAuthInvalidToken = &AError{Type: "auth", Msg: "invalid token", Status: http.StatusUnauthorized}

	ErrRequestBadParam    = &AError{Type: "request", Msg: "bad request", Status: http.StatusBadRequest}
	ErrRequestInvalidCred = &AError{Type: "request", Msg: "invalid credentials", Status: http.StatusBadRequest}
)
