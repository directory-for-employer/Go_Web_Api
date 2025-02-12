package auth

import (
	"go/web-api/configs"
	"go/web-api/internal/user"
	"go/web-api/pkg/jwt"
	"go/web-api/pkg/req"
	"go/web-api/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}
type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/reg", handler.Registration())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		_, err = handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			res.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJwt(handler.Config.Auth.Secret).Create(jwt.JWTData{
			Email: body.Email,
		})
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := LoginResponse{
			Token: token,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Registration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		userData := &user.User{
			Email:    body.Email,
			Name:     body.Name,
			Password: body.Password,
		}
		_, err = handler.AuthService.Register(userData)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := jwt.NewJwt(handler.Config.Auth.Secret).Create(jwt.JWTData{
			Email: body.Email,
		})
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := RegisterResponse{
			Token: token,
		}
		res.Json(w, data, http.StatusOK)
	}
}
