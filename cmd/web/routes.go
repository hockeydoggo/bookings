package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hockeydoggo/bookings/pkg/config"
	"github.com/hockeydoggo/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
