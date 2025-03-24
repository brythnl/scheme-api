package handler

import (
	"net/http"

	"github.com/brythnl/scheme-api/internal/model"
	"github.com/brythnl/scheme-api/internal/service"
)

func Login(s service.AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginData, err := decode[model.UserLogin](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := s.Login(r.Context(), loginData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		encode(w, http.StatusOK, token)
	})
}

func Register(s service.AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData, err := decode[model.UserCreate](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := s.Register(r.Context(), userData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		encode(w, http.StatusCreated, user)
	})
}
