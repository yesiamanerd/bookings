package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yesiamanerd/bookings/pkg/config"
	"github.com/yesiamanerd/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	// Mux using the type of chi router
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
