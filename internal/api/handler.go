package api

import (
	"net/http"
)

type API struct {
	authService auth
}

func NewAPI(a auth) API {
	return API{authService: a}
}

func (a API) Auth(writer http.ResponseWriter, request *http.Request) {
	login := request.FormValue("login")
	pass := request.FormValue("password")

	if login == "" {
		http.Error(writer, "empty login", http.StatusBadRequest)
		return
	}
	if login == "" {
		http.Error(writer, "empty password ", http.StatusBadRequest)
		return
	}

	state, err := a.authService.Auth(request.Context(), login, pass)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	if !state.Valid {
		http.Error(writer, "invalid state", http.StatusUnauthorized)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:  "auth",
		Value: "good",
	})
}

func NewAuthHandler(a auth) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		login := request.FormValue("login")
		pass := request.FormValue("password")

		if login == "" {
			http.Error(writer, "empty login", http.StatusBadRequest)
			return
		}
		if login == "" {
			http.Error(writer, "empty password ", http.StatusBadRequest)
			return
		}

		state, err := a.Auth(request.Context(), login, pass)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusUnauthorized)
			return
		}

		if !state.Valid {
			http.Error(writer, "invalid state", http.StatusUnauthorized)
			return
		}

		http.SetCookie(writer, &http.Cookie{
			Name:  "auth",
			Value: "good",
		})
	}
}
