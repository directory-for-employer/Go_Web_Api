package user

import (
	"go/web-api/pkg/req"
	"go/web-api/pkg/res"
	"net/http"
)

type UserHandlerDeps struct {
	UserRepository *UserRepository
}
type UserHandler struct {
	UserRepository *UserRepository
}

func NewUserHandler(router *http.ServeMux, deps UserHandlerDeps) {
	handler := &UserHandler{
		UserRepository: deps.UserRepository,
	}
	router.HandleFunc("POST /user/create", handler.Create())
	//router.Handle("POST /user/login", middleware.IsAuthed(handler.Login()))
}

func (handler *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[UserCreateRequest](&w, r)
		newUser := &User{
			Email:    body.Email,
			Password: body.Password,
			Name:     body.Name,
		}
		if err != nil {
			return
		}
		dataUser, err := handler.UserRepository.Create(newUser)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//	return
		//}
		res.Json(w, dataUser, http.StatusCreated)
	}
}

//func (handler *UserHandler) Login() http.HandlerFunc {
//
//}
