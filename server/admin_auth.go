package server

import (
	"net/http"

	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/asaskevich/govalidator"
	"github.com/panjiesw/apimocker/db"
	"github.com/panjiesw/apimocker/errs"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"golang.org/x/crypto/bcrypt"
)

type LoginForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Remember bool   `json:"remember" valid:"-"`
}

func (s *Server) adminAuthRoute() chi.Router {
	r := chi.NewRouter()
	r.Post("/login", s.AdminAuthLogin)
	return r
}

func (s *Server) AdminAuthLogin(w http.ResponseWriter, r *http.Request) {
	var form LoginForm
	if err := render.Bind(r.Body, &form); err != nil {
		RenderAError(w, r, errs.ErrRequestBadParam)
		return
	}

	if valid, err := govalidator.ValidateStruct(form); err != nil || !valid {
		RenderAError(w, r, errs.ErrRequestBadParam)
		return
	}

	var user *db.User
	var err *errs.AError
	if govalidator.IsEmail(form.Login) {
		user, err = s.DS.UserGetByEmail(form.Login)
	} else {
		user, err = s.DS.UserGetByUsername(form.Login)
	}

	if err != nil {
		RenderAError(w, r, err)
		return
	}

	if err = validatePassword(user.Password, form.Password); err != nil {
		RenderAError(w, r, err)
		return
	}

	token, err := s.generateToken(user)
	if err != nil {
		RenderAError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]string{"token": token})
}

func (s *Server) generateToken(u *db.User) (string, *errs.AError) {
	now := time.Now().Add(time.Duration(100) * time.Nanosecond)

	claims := jws.Claims{}
	claims.SetNotBefore(now)
	claims.SetExpiration(now.Add(time.Duration(24) * 7 * time.Hour))
	claims.SetIssuedAt(now)
	claims.SetIssuer("apimocker")
	claims.SetAudience("http://localhost:3000")

	jwt := jws.NewJWT(claims, crypto.SigningMethodHS512)
	token, err := jwt.Serialize("somekey")
	if err != nil {
		return "", errs.New("auth", "failed to serialize token", 500)
	}
	return string(token), nil
}

func validatePassword(hashed, plain string) *errs.AError {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return errs.ErrRequestInvalidCred
	}
	return nil
}
