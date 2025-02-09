package auth

import (
	"fmt"
	"go/web-api/configs"
	"go/web-api/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/reg", handler.Registration())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Login")
		data := LoginResponse{
			Token: handler.Auth.Secret,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Registration() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Registration")
	}
}
