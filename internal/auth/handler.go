package auth

import (
	"errors"
	"net/http"
	"project/configs"
	"project/pkg/jwt"
	"project/pkg/req"
	"project/pkg/res"

	"gorm.io/gorm"
)

type AuthHandlerDeps struct {
	*configs.Config
	AuthRepository *AuthRepository
}

type Handler struct {
	*configs.Config
	AuthRepository *AuthRepository
}

func NewHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &Handler{
		Config:         deps.Config,
		AuthRepository: deps.AuthRepository,
	}
	router.HandleFunc("POST /auth/phone", handler.Phone())
	router.HandleFunc("POST /auth/phonecode", handler.PhoneCode())
}

func (handler *Handler) Phone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[PhoneRequest](&w, r)
		if err != nil {
			return
		}
		err = handler.AuthRepository.GetPhone(body.Phone)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newPhone := NewPhone(body.Phone)
			createdPhone, err := handler.AuthRepository.CreatePhone(newPhone)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			resp := PhoneResponse{
				SessionId: createdPhone.SessionId,
			}
			res.Json(w, resp, http.StatusCreated)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updatedPhone := NewPhone(body.Phone)
		newUpdatedPhone, err := handler.AuthRepository.UpatePhone(updatedPhone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp := PhoneResponse{
			SessionId: newUpdatedPhone.SessionId,
		}
		res.Json(w, resp, http.StatusCreated)
	}
}

func (handler *Handler) PhoneCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[PhoneRequestWithCode](&w, r)
		if err != nil {
			return
		}
		phone, err := handler.AuthRepository.GetPhoneByCode(body.SessionId, body.Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{
			Phone: phone.Phone,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := &PhoneResponseToken{
			JWT: token,
		}

		res.Json(w, resp, http.StatusCreated)
	}
}
