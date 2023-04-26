package main

import (
	"github.com/gandra/helloworld-web-go-trevor-sawler/config"
	"github.com/gandra/helloworld-web-go-trevor-sawler/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Use(middleware.Recoverer)

	return mux
}
