package main

import (
	"net/http"

	"go-atch-template/internal/api"
	"go-atch-template/internal/service/auth"
	"go-atch-template/internal/storage/user"
)

func main() {
	cfg := parseConfig()

	userStorage := user.New(cfg.PostgresDB)

	authService := auth.New(userStorage)

	a := api.NewAPI(authService)

	// authHandler := api.NewAuthHandler(authService)

	_ = http.ListenAndServe(cfg.ListenAddr, http.HandlerFunc(a.Auth))

}

func parseConfig() config {
	return config{}
}

type config struct {
	PostgresDB string // flag, env, file
	ListenAddr string // flag, env, file
}
