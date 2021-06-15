package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kranthy09/go-course/pkg/config"
	"github.com/kranthy09/go-course/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(Nosurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
