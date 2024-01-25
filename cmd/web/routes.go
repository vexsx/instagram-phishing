package main

import (
	"github.com/go-chi/chi/v5"
	"insta/pkg/config"
	"insta/pkg/handler"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handler.Repo.Index)

	return mux
}
