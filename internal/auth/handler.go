package auth

import (
	"net/http"
	"project/configs"
	"project/pkg/req"
	"project/pkg/res"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type Handler struct {
	*configs.Config
}

func NewHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &Handler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		resp := LoginResponse{
			Token: "",
		}

		res.Json(w, resp, http.StatusCreated)
	}
}

func (handler *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		resp := RegisterResponse{
			Token: "",
		}

		res.Json(w, resp, http.StatusCreated)
	}
}
