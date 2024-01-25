package main

import (
	"github.com/go-chi/chi/v5"
	"insta/pkg/config"
	"insta/pkg/handler"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//middleware
	mux.Use(NoSurf)

	//Serve static files from the "static" directory
	staticFileDir := http.Dir("./static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(staticFileDir))
	mux.Handle("/static/*", staticFileServer)

	// Define your routes using chi router
	//mux.HandleFunc("/", handler.Repo.Index)
	mux.Get("/", handler.Repo.Index)

	return mux
}
