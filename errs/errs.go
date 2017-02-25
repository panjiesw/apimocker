package errs

import (
	"fmt"
)

type AError struct {
	Type   string `json:"type"`
	Msg    string `json:"message"`
	Status int    `json:"-"`
}

func (e *AError) Error() string {
	return fmt.Sprintf("%s - %s", e.Type, e.Msg)
}

var (
	// ErrUsernameExists is returned when trying to signup using existing username.
	ErrUsernameExists    = &AError{Type: "db", Msg: "username already exists", Status: 400}
	ErrUsernameNotExists = &AError{Type: "db", Msg: "username doesn't exist", Status: 404}
	// ErrEmailExists is returned when trying to signup using existing email.
	ErrEmailExists    = &AError{Type: "db", Msg: "email already exists", Status: 400}
	ErrEmailNotExists = &AError{Type: "db", Msg: "email doesn't exist", Status: 404}
	ErrIDNotExists    = &AError{Type: "db", Msg: "id doesn't exist", Status: 404}

	ErrDBUnkown = &AError{Type: "db", Msg: "internal error", Status: 500}

	ErrAuthNoToken      = &AError{Type: "auth", Msg: "no token", Status: 401}
	ErrAuthInvalidToken = &AError{Type: "auth", Msg: "invalid token", Status: 401}
)
