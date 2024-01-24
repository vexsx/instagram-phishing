package main

import (
	"github.com/bmizerany/pat"
	"insta/pkg/config"
	"insta/pkg/handler"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Repo.Index))
	return mux
}
