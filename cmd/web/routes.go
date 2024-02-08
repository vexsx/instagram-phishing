package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"insta/pkg/config"
	"insta/pkg/handler"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//middleware
	mux.Use(middleware.Logger)

	mux.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	//Serve static files from the "static" directory
	staticFileDir := http.Dir("./static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(staticFileDir))
	mux.Handle("/static/*", staticFileServer)

	// Define your routes using chi router
	//mux.HandleFunc("/", handler.Repo.Index)
	mux.Get("/", handler.Repo.Index)
	mux.Post("/v1/Login", handler.Repo.LoginHandler)
	mux.Post("/v2/Login", handler.Repo.LoginHandlerFilter)

	return mux
}
